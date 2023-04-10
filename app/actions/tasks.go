package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"todos/app/models"
)

// TasksResource is the resource for the Task model
type TasksResource struct {
	buffalo.Resource
}

// List gets all Tasks. This function is mapped to the path
// GET /tasks
func (v TasksResource) List(c buffalo.Context) error {

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	tasks := models.Tasks{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".

	q := tx.PaginateFromParams(c.Params())

	if c.Param("section") == "complete" {
		q.Where("complete")
	} else if c.Param("section") == "incomplete" {
		q.Where("NOT complete")
	} else {
		return c.Error(http.StatusNotFound, fmt.Errorf("could not find %s", c.Request().URL))
	}

	// Retrieve all Tasks from the DB
	if err := q.All(&tasks); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("tasks", tasks)

	count, err := CountTasks(tx)
	if err != nil {
		fmt.Println(err)

	}
	c.Set("count", count)

	return c.Render(http.StatusOK, r.HTML("/tasks/index.plush.html"))
}

// Show gets the data for one Task. This function is mapped to
// the path GET /tasks/{task_id}
func (v TasksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Task
	task := &models.Task{}

	// To find the Task the parameter task_id is used.
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("/tasks/show.plush.html"))
}

// New renders the form for creating a new Task.
// This function is mapped to the path GET /tasks/new
func (v TasksResource) New(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	task := &models.Task{}

	count, err := CountTasks(tx)
	if err != nil {
		fmt.Println(err)

	}
	c.Set("count", count)
	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("/tasks/new.plush.html"))
}

// Create adds a Task to the DB. This function is mapped to the
// path POST /tasks
func (v TasksResource) Create(c buffalo.Context) error {
	// Allocate an empty Task
	task := &models.Task{}

	// Bind task to the html form elements
	if err := c.Bind(task); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	count, err := CountTasks(tx)
	if err != nil {
		fmt.Println(err)

	}
	c.Set("count", count)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(task)
	if err != nil {
		return err
	}

	if verrs.HasAny() {

		fmt.Println(task)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.

		c.Set("task", task)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/tasks/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "task.created.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/tasks/incomplete")

}

// Edit renders a edit form for a Task. This function is
// mapped to the path GET /tasks/{task_id}/edit
func (v TasksResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Task
	task := &models.Task{}

	if err := tx.Find(task, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	count, err := CountTasks(tx.Q().Connection)
	if err != nil {
		fmt.Println(err)

	}
	c.Set("count", count)
	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("/tasks/edit.plush.html"))
}

// Update changes a Task in the DB. This function is mapped to
// the path PUT /tasks/{task_id}
func (v TasksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Task
	task := &models.Task{}

	if err := tx.Find(task, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Task to the html form elements
	if err := c.Bind(task); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(task)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("task", task)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/tasks/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "task.updated.success")

	if task.Complete {
		return c.Redirect(http.StatusSeeOther, "/tasks/complete")
	}

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/tasks/incomplete")
}

// Destroy deletes a Task from the DB. This function is mapped
// to the path DELETE /tasks/{task_id}
func (v TasksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Task
	task := &models.Task{}

	// To find the Task the parameter task_id is used.
	if err := tx.Find(task, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(task); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "task.destroyed.success")

	if task.Complete {
		return c.Redirect(http.StatusSeeOther, "/tasks/complete")
	}
	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "/tasks/incomplete")
}

func Complete(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	task := &models.Task{}

	// To find the Task the parameter task_id is used.

	if err := tx.Find(task, c.Param("id")); err != nil {
		return fmt.Errorf("no existe ese id buscado")
	}

	task.Complete = !task.Complete
	err := tx.Update(task)
	if err != nil {
		return err
	}

	if task.Complete {
		return c.Redirect(http.StatusSeeOther, "/tasks/complete")
	}

	return c.Redirect(http.StatusSeeOther, "/tasks/incomplete")

}

func CountTasks(tx *pop.Connection) (int, error) {
	task := &models.Task{}

	count, err := tx.Where("NOT complete").Count(task)
	if err != nil {
		return count, err
	}

	return count, nil
}

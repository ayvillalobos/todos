package app

import (
	"net/http"

	"todos/app/actions"
	"todos/app/actions/home"
	"todos/app/middleware"
	"todos/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)

	taskResource := actions.TasksResource{}

	//root.Resource("/tasks", taskResource)
	//root.GET("/tasks/", taskResource.List)
	root.GET("/tasks/{section}", taskResource.List)
	root.DELETE("/tasks/{id}", taskResource.Destroy)
	root.PUT("/tasks/{id}", taskResource.Update)
	root.GET("/task/new", taskResource.New)
	root.GET("/task/edit", taskResource.Edit)
	root.POST("/tasks", taskResource.Create)
	root.PATCH("/task/{id}", actions.Complete).Name("completeTaskPath")

	root.ServeFiles("/", http.FS(public.FS()))

}

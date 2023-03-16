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
	root.Resource("/tasks", actions.TasksResource{})
	root.PATCH("/tasks/{task_id}/", actions.Complete)

	root.ServeFiles("/", http.FS(public.FS()))
}

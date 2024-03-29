package home

import (
	"net/http"
	"todos/app/render"

	"github.com/gobuffalo/buffalo"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Index(c buffalo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/tasks/incomplete")
}

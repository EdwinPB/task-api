package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// TasksList default implementation.
func TasksList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("tasks/list.html"))
}

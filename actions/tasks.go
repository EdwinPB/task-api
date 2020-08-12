package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// TasksList default implementation.
func TasksList(c buffalo.Context) error {
	fmt.Println("---_>")
	return c.Render(http.StatusOK, r.HTML("tasks/list.html"))
}

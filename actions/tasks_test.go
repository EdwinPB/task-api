package actions

import (
	"net/http"
)

func (as *ActionSuite) Test_Tasks_List() {
	res := as.HTML("/tasks/list").Get()
	as.Equal(http.StatusOK, res.Code)
}

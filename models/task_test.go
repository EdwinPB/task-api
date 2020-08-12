package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Task() {
	task := Task{}
	verrs, err := task.Validate(ms.DB)
	ms.Error(verrs)
	ms.NoError(err)

	task = Task{
		ID:             uuid.Must(uuid.NewV4()),
		Description:    "New task",
		ExecutorName:   "Edwin",
		RequesterName:  "Larry",
		Completed:      true,
		CompletionDate: nulls.Time{},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	verrs, err = task.Validate(ms.DB)
	ms.Empty(verrs.Error())
	ms.NoError(err)

	err = ms.DB.Create(&task)
	ms.NoError(err)
}

func (ms *ModelSuite) Test_Task_Storage() {
	ts := TasksStorage{}
	ts = ts.Add(Task{})
	ms.Equal(1, len(ts))
	ts = ts.Add(Task{})
	ms.Equal(2, len(ts))
}

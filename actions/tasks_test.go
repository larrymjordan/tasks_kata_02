package actions

import (
	"net/http"
	"tasks/models"

	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_Tasks_List() {
	taskStorage = &models.Tasks{}

	resp := as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, resp.Code)

	var tasks []models.Task
	resp.Bind(&tasks)

	as.Len(tasks, 0)

	taskStorage.Create(models.Task{
		Description: "Do this",
		IsDone:      true,
	})

	resp = as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, resp.Code)

	resp.Bind(&tasks)

	as.Len(tasks, 1)
	as.Equal(taskStorage.List()[0].Description, tasks[0].Description)
	as.Equal(taskStorage.List()[0].IsDone, tasks[0].IsDone)
}

func (as *ActionSuite) Test_Tasks_Create() {
	taskStorage = &models.Tasks{}
	resp := as.JSON("/tasks").Post(models.Task{
		Description: "Do this",
		IsDone:      true,
	})

	as.Equal(http.StatusCreated, resp.Code)

	createdTask := models.Task{}
	resp.Bind(&createdTask)

	as.NotEqual(uuid.Nil, createdTask.ID)
	as.Equal("Do this", createdTask.Description)

	as.Len(taskStorage.List(), 1)
}

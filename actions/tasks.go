package actions

import (
	"net/http"
	"tasks/models"

	"github.com/gobuffalo/buffalo"
)

var taskStorage models.TaskStorage = &models.DBTaskStorage{models.DB}

// TasksList default implementation.
func TasksList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(taskStorage.List()))
}

// TasksCreate default implementation.
func TasksCreate(c buffalo.Context) error {
	task := &models.Task{}
	c.Bind(task)

	taskStorage.Create(task)
	return c.Render(http.StatusCreated, r.JSON(task))
}

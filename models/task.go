package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Task is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	IsDone      bool      `json:"is_done" db:"is_done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

type TaskStorage interface {
	List() Tasks
	Create(Task)
}

// Tasks is not required by pop and may be deleted
type Tasks []Task

func (t *Tasks) List() Tasks {
	return *t
}

func (t *Tasks) Create(task Task) {
	*t = append(*t, task)
}

// String is not required by pop and may be deleted
func (t Tasks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

type DBTaskStorage struct {
	Tx *pop.Connection
}

func (t *DBTaskStorage) List() Tasks {
	tasks := Tasks{}
	t.Tx.All(&tasks)
	return tasks
}

func (t *DBTaskStorage) Create(task Task) {
	t.Tx.Create(&task)
}

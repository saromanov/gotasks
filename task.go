package gotasks

import (
	"github.com/google/uuid"
)

// Task defines instance for task
type Task struct {
	ID     string
	Name   string
	Method func(*Entry) error
}

// NewTask provides create of new task
func NewTask(name string, method func(*Entry) error) *Task {
	return &Task{
		ID:     uuid.New().String(),
		Name:   name,
		Method: method,
	}
}

// GetID returns id of the task
func (t *Task) GetID() string {
	return t.ID
}

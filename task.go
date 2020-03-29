package gotasks

import (
	"context"
	"github.com/google/uuid"
)
// Task defines instance for task
type Task struct {
	ID string
	Name string
	Method func(context.Context) error
}

// NewTask provides create of new task
func NewTask(name string, method func(context.Context) error) *Task {
	return &Task{
		ID: uuid.New().String(),
		Name: name,
		Method: method,
	}
}
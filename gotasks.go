package gotasks

import (
	"context"
	"errors"
	"sync"
)

var errTaskNotFound = errors.New("task is not found")

// GoTasks provides implementation of tasks
type GoTasks struct {
	mu    *sync.RWMutex
	tasks map[string]*Task
}

// New provides creating of the tasks instance
func New() *GoTasks {
	return &GoTasks{
		mu:    &sync.RWMutex{},
		tasks: make(map[string]*Task),
	}
}

// Add provides addding of the task
func (g *GoTasks) Add(name string, f func(context.Context) error) (string, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	g.tasks[name] = NewTask(name, f)
	return "", nil
}

// Exec provides execution of task
func (g *GoTasks) Exec(name string) error {
	task, ok := g.tasks[name]
	if !ok {
		return errTaskNotFound
	}
	go task.Method(context.Background())
	return nil
}

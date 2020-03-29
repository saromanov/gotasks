package gotasks

import (
	"context"
	"sync"
)

// GoTasks provides implementation of tasks
type GoTasks struct {
	mu    *sync.RWMutex
	tasks map[string]func(context.Context) error
}

// New provides creating of the tasks instance
func New() *GoTasks {
	return &GoTasks{
		mu:    &sync.RWMutex{},
		tasks: make(map[string]func(context.Context) error),
	}
}

// Add provides addding of the task
func (g *GoTasks) Add(name string, f func(context.Context) error) (string, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	g.tasks[name] = f
	return "", nil
}

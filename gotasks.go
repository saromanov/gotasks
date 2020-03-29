package gotasks

import "context"

// GoTasks provides implementation of tasks
type GoTasks struct {
}

// Add provides addding of the task
func (g *GoTasks) Add(f func(context.Context) error) error {
	return nil
}

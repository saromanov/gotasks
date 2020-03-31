package gotasks

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var errTaskNotFound = errors.New("task is not found")

// ExecOption defines options for execution of the task
type ExecOption func(*Option)

type Option struct {
	timeout time.Duration
}

// WithTimeout defines option with specific timeput
func WithTimeout(d time.Duration) ExecOption {
	return func(opt *Option) {
		opt.timeout = d
	}
}

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
func (g *GoTasks) Add(name string, f func(*Entry) error) (string, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	t := NewTask(name, f)
	g.tasks[name] = t
	return t.GetID(), nil
}

// Exec provides execution of task
func (g *GoTasks) Exec(name string, opt ...ExecOption) error {
	task, ok := g.tasks[name]
	if !ok {
		return errTaskNotFound
	}

	options := &Option{}
	for _, o := range opt {
		o(options)
	}
	singleExec(task, options)
	return nil
}

func singleExec(tas *Task, opt *Option) {
	go func(t *Task) {
		t.Method(&Entry{})
	}(tas)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	if opt.timeout != 0*time.Second {
		ctx, cancel = context.WithTimeout(ctx, opt.timeout)
	}
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	}
}

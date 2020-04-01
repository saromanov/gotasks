package gotasks

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var errTaskNotFound = errors.New("task is not found")

// ExecOption defines options for execution of the task
type ExecOption func(*Option)

type Option struct {
	timeout       time.Duration
	cancelFunc    func(*Entry)
	numGoroutines int
}

// WithTimeout defines option with specific timeput
func WithTimeout(d time.Duration, cancelFunc func(*Entry)) ExecOption {
	return func(opt *Option) {
		opt.timeout = d
		opt.cancelFunc = cancelFunc
	}
}

// WithPool option defines setting of goroutine pool
// to the execution
func WithPool(num int) ExecOption {
	return func(opt *Option) {
		opt.numGoroutines = num
	}
}

// GoTasks provides implementation of tasks
type GoTasks struct {
	mu      *sync.RWMutex
	tasks   map[string]*Task
	running int32
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

// Running returns current runnign tasks
func (g *GoTasks) Running() int32 {
	return atomic.LoadInt32(&g.running)
}

// Exec provides execution of task
func (g *GoTasks) Exec(name string, opt ...ExecOption) error {
	task, ok := g.tasks[name]
	if !ok {
		return errTaskNotFound
	}

	options := func(o []ExecOption) *Option {
		opts := &Option{}
		for _, o := range opt {
			o(opts)
		}
		return opts
	}(opt)
	ctx, cancel := makeContext(options)
	defer cancel()
	singleExec(ctx, task, options)
	return nil
}

func makeContext(opt *Option) (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	if opt.timeout == 0*time.Second {
		return ctx, cancel
	}
	return context.WithTimeout(ctx, opt.timeout)
}

func singleExec(ctx context.Context, tas *Task, opt *Option) {
	go func(t *Task) {
		t.Method(&Entry{})
	}(tas)
	select {
	case <-ctx.Done():
		if opt.cancelFunc != nil {
			opt.cancelFunc(&Entry{Ctx: ctx})
			return
		}
		fmt.Println(ctx.Err())
		return
	}
}

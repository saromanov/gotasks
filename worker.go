package gotasks

import "fmt"

type workRequest struct {
	jobChan       chan<- Task
	retChan       <-chan interface{}
	payload       interface{}
	interruptFunc func()
}

// worker implements execution of tasks on pool
type worker struct {
	closeChan  chan struct{}
	closedChan chan struct{}
	reqChan    chan<- workRequest
	task       *Task
}

func newWorker(t *Task, reqChan chan<- workRequest) *worker {
	return &worker{
		task:       t,
		closeChan:  make(chan struct{}),
		closedChan: make(chan struct{}),
		reqChan:    reqChan,
	}
}

func (w *worker) run() {
	jobChan, retChan := make(chan Task), make(<-chan interface{})
	defer func() {
		//close(retChan)
		close(w.closedChan)
	}()
	for {
		select {
		case w.reqChan <- workRequest{
			jobChan: jobChan,
			retChan: retChan,
		}:
			fmt.Println("AAA")
			select {
			case data := <-jobChan:
				fmt.Println("WORK REQYST: ", data.Method(&Entry{}))

			}
		case <-w.closeChan:
			return
		}
	}
}

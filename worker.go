package gotasks

type workRequest struct {
	jobChan       chan<- interface{}
	retChan       <-chan interface{}
	payload       interface{}
	interruptFunc func()
}

// worker implements execution of tasks on pool
type worker struct {
	closeChan  chan struct{}
	closedChan chan struct{}
	reqChan    chan workRequest
	task       *Task
}

func newWorker(t *Task) *worker {
	return &worker{
		task:       t,
		closeChan:  make(chan struct{}),
		closedChan: make(chan struct{}),
		reqChan:    make(chan workRequest),
	}
}

func (w *worker) run() {
	_, retChan := make(chan interface{}), make(chan interface{})
	defer func() {
		close(retChan)
		close(w.closedChan)
	}()

	for {
		select {
		case _ = <-w.reqChan:
			w.task.Method(&Entry{})
		case <-w.closeChan:
			return
		}
	}
}

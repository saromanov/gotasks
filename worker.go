package gotasks


type workRequest struct {
	jobChan chan<- interface{}
	retChan <-chan interface{}
	interruptFunc func()
}

// worker implements execution of tasks on pool
type worker struct {
	closeChan chan struct{}
	closedChan chan struct{}
	reqChan chan<- workRequest
	task *Task
}

func newWorker(t *Task) *worker {
	return &worker{
		task:t,
		closeChan:     make(chan struct{}),
		closedChan:    make(chan struct{}),
	}
}

func (w *worker) run(){
	jobChan, retChan := make(chan interface{}), make(chan interface{})
	defer func() {
		close(retChan)
		close(w.closedChan)
	}()

	for {
		select {
		case w.reqChan <- workRequest{
			jobChan:       jobChan,
			retChan:       retChan,
		}:
			select {
			case _ = <-jobChan:
				result := w.task.Method(&Entry{})
				select {
				case retChan <- result:
				}
			}
		case <-w.closeChan:
			return
		}
	}
}
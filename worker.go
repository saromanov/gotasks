package gotasks


// worker implements execution of tasks on pool
type worker struct {
	closeChan chan struct{}
	closedChan chan struct{}
	task *Task
}

func newWorker(t *Task) *worker {
	return &worker{
		t:t,
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
			interruptFunc: w.interrupt,
		}:
			select {
			case payload := <-jobChan:
				result := w.worker.Process(payload)
				select {
				case retChan <- result:
				case <-w.interruptChan:
					w.interruptChan = make(chan struct{})
				}
			case _, _ = <-w.interruptChan:
				w.interruptChan = make(chan struct{})
			}
		case <-w.closeChan:
			return
		}
	}
}
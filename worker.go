package gotasks


// worker implements execution of tasks on pool
type worker struct {
	closeChan chan struct{}
	closedChan chan struct{}
}

func newWorker() *worker {
	return &worker{
		closeChan:     make(chan struct{}),
		closedChan:    make(chan struct{}),
	}
}
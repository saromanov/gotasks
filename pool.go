package gotasks

// poolExec provides execution of the task on the pool
// of goroutines
func poolExec(num int, t *Task) error {
	w := newWorker(t)
	reqChan := make(chan workRequest)
	request := <-reqChan
	request.jobChan <- "a"
	w.run()
	return nil
}

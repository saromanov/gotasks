package gotasks

// poolExec provides execution of the task on the pool
// of goroutines
func poolExec(num int, t *Task) error {
	workReq := make(chan workRequest)
	workers := make([]*worker, num)
	for i := 0; i < num; i++ {
		workers[i] = newWorker(t, workReq)
		go workers[i].run()
		request, _ := <-workReq
		request.jobChan <- *t
	}
	return nil
}

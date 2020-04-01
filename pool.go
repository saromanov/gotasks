package gotasks

// poolExec provides execution of the task on the pool
// of goroutines
func poolExec(num int, t *Task) error {
	w := newWorker(t)
	w.run()
	return nil	
}
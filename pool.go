package gotasks

import "fmt"

// poolExec provides execution of the task on the pool
// of goroutines
func poolExec(num int, t *Task) error {
	w := newWorker(t)
	fmt.Println("FAIN")
	go w.run()
	w.reqChan <- workRequest{}
	fmt.Println("FAIN")
	return nil
}

package main

import (
	"context"

	"github.com/saromanov/gotasks"
)

func add(ctx context.Context) int {
	a := ctx.Value("a").(int)
	b := ctx.Value("b").(int)
	return a + b
}
func main() {
	tasks := gotasks.GoTasks{}
	tasks.Add(add)
}

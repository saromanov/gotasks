package main

import (
	"context"

	"github.com/saromanov/gotasks"
)

func add(ctx context.Context) error {
	a := ctx.Value("a").(int)
	b := ctx.Value("b").(int)
	context.WithValue(ctx, "add_func", a+b)
	return nil
}

func sub(ctx context.Context) error {
	a := ctx.Value("a").(int)
	b := ctx.Value("b").(int)
	context.WithValue(ctx, "sub_func", a-b)
	return nil
}

func main() {
	tasks := gotasks.New()
	tasks.Add("add", add)
}

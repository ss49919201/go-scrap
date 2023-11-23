package main

import (
	"context"
	"fmt"
	"time"
)

func main() {}

func a() {
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// 親コンテキストへは伝播しないので、`fatal error: all goroutines are asleep - deadlock!`
	<-ctx.Done()
	if ctx.Err() != nil {
		fmt.Println(ctx.Err())
	}
}

func b() {
	ctx := context.Background()
	var child context.Context
	var cancel context.CancelFunc
	child, cancel = context.WithTimeout(ctx, time.Second)
	defer cancel()

	<-child.Done()
	if child.Err() != nil {
		fmt.Println(child.Err()) // context deadline exceeded

		// 親はキャンセルされていない
		fmt.Println(ctx.Done())       // <nil>
		fmt.Println(ctx.Err() == nil) // true
	}
}

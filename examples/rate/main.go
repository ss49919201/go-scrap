package main

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	printInterval1Sec()
}

func printInterval1Sec() {
	ctx := context.Background()
	l := rate.NewLimiter(rate.Every(time.Duration(1000000000)), 1)
	for i := 0; i < 10; i++ {
		if err := l.Wait(ctx); err != nil {
			panic(err)
		}
		println(i)
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	c()
}

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

func c() {
	// debug: curl -Ssi "http://localhost:8080/ping/"
	http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
		child, cancel := context.WithTimeoutCause(r.Context(), time.Second*5, errors.New("cause error"))
		defer cancel()
		r = r.WithContext(child)

		done := make(chan struct{})
		go func() {
			time.Sleep(time.Second * 10)
			done <- struct{}{}
		}()

		go func() {
			time.Sleep(time.Second * 3)
			cancel()
		}()

		select {
		case <-child.Done():
			if errors.Is(child.Err(), context.Canceled) {
				fmt.Println("cancled")
				fmt.Println("child.Err()", child.Err())
				fmt.Println("context.Cause(child)", context.Cause(child))
			} else if errors.Is(child.Err(), context.DeadlineExceeded) {
				fmt.Println("deadline exceeded")
				fmt.Println("child.Err()", child.Err())
				fmt.Println("context.Cause(child)", context.Cause(child))
			} else {
				fmt.Println("something")
				fmt.Println("child.Err()", child.Err())
				fmt.Println("context.Cause(child)", context.Cause(child))
			}
		case <-done:
			fmt.Println("done process")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
		if id := strings.TrimPrefix(r.URL.Path, "/ping/"); id != "" {
			w.Write([]byte(fmt.Sprintf("id = %s\n", id)))
		}
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

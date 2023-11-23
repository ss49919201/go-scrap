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
	run()
}

func run() {
	// debug: curl -Ssi "http://localhost:8080/ping/"
	http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
		child, cancelChild := context.WithCancelCause(r.Context())

		grandChild, cancel := context.WithTimeoutCause(child, time.Second*10, errors.New("cause error"))
		defer cancel()
		r = r.WithContext(grandChild)

		done := make(chan struct{})
		go func() {
			time.Sleep(time.Second * 10)
			done <- struct{}{}
		}()

		go func() {
			time.Sleep(time.Microsecond)
			cancelChild(errors.New("cancel child"))
		}()

		select {
		case <-grandChild.Done():
			if errors.Is(grandChild.Err(), context.Canceled) {
				fmt.Println("cancled")
				fmt.Println("grandChild.Err()", grandChild.Err())
				fmt.Println("context.Cause(grandChild)", context.Cause(grandChild))
			} else if errors.Is(grandChild.Err(), context.DeadlineExceeded) {
				fmt.Println("deadline exceeded")
				fmt.Println("grandChild.Err()", grandChild.Err())
				fmt.Println("context.Cause(grandChild)", context.Cause(grandChild))
			} else {
				fmt.Println("something")
				fmt.Println("grandChild.Err()", grandChild.Err())
				fmt.Println("context.Cause(grandChild)", context.Cause(grandChild))
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

	srv := http.Server{
		Addr:              "127.0.0.1:8080",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	srv.ListenAndServe()
}

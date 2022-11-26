package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var (
	hostname string
	port     string
	address  string
)

func setenv() {
	hostname = os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "127.0.0.1"
	}
	port = os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	address = net.JoinHostPort(hostname, port)
}

func corsMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// CORS の許可
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTION")
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func loggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		log.Printf("[%v] ", r.Method)
	}
}

// https://golang.org/doc/articles/wiki/#tmp_1
func Start() {
	setenv()

	srvMux := http.NewServeMux()

	// example
	srvMux.HandleFunc(
		"/example",
		loggerMiddleware(
			corsMiddleware(
				func(w http.ResponseWriter, r *http.Request) {
					b, err := json.Marshal(
						struct {
							String string
							Int    int
							Bool   bool
						}{
							String: "example",
							Int:    1,
							Bool:   true,
						},
					)
					if err != nil {
						panic(err)
					}
					io.WriteString(w, string(b))
				},
			),
		),
	)

	// ping
	srvMux.HandleFunc(
		"/ping",
		loggerMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				b, err := json.Marshal("{ping:pong}")
				if err != nil {
					panic(err)
				}
				io.WriteString(w, string(b))
			},
		),
	)

	// hello
	srvMux.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, fmt.Sprintf("Hello! %s.", r.UserAgent()))
		},
	)

	srv := &http.Server{
		Addr:    address,
		Handler: srvMux,
	}

	idleConnsClosed := make(chan struct{}, 1)
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	fmt.Printf("listen server address: %s\n", address)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	fmt.Println("Server closed.")
}

func Request() {
	http.Get("http://localhost:12345/example")
}

func main() {
	if os.Args[1] == "1" {
		Start()
	} else if os.Args[1] == "2" {
		Request()
	} else {
		fmt.Println("invalid argument")
	}
}

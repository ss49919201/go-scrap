package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	hostname string
	port     string
	address  string
)

func setenv() {
	hostname = os.Getenv("HOSTNAME")
	port = os.Getenv("PORT")
	address = fmt.Sprintf("%s:%s", hostname, port)
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

	// example
	http.HandleFunc(
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
	http.HandleFunc(
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
	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, fmt.Sprintf("Hello! %s.", r.UserAgent()))
		},
	)

	// Only localhost
	fmt.Printf("listen server address: %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

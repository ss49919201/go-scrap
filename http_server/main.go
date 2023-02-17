package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// debug: curl -Ssi "http://localhost:8080/ping/"
	http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
		if id := strings.TrimPrefix(r.URL.Path, "/ping/"); id != "" {
			w.Write([]byte(fmt.Sprintf("id = %s\n", id)))
		}
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

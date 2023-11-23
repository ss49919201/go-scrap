package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	run()
}

func run() {
	client := new(http.Client)
	client.Timeout = 1 * time.Second
	resp, err := client.Get("http://localhost:8080/ping/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("status code = %d\n", resp.StatusCode)
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("body = %s\n", b)
}

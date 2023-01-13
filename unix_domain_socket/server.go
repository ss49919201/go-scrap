package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

var address string

func init() {
	wd, _ := os.Getwd()
	address = path.Join(wd, "temp", "example")
	os.Remove(address)
}

func main() {
	listner, err := net.Listen("unix", address)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	defer func() {
		err := os.Remove(address)
		if err != nil {
			log.Default().Println(err)
		}
	}()
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		req, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		body, _ := io.ReadAll(req.Body)
		os.Stdout.WriteString(fmt.Sprintf("read request: %v\n", string(body)))
		(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("This is response")),
		}).Write(conn)
		conn.Close()
	}
}

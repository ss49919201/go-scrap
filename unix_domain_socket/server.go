package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	wd, _ := os.Getwd()
	listner, err := net.Listen("unix", path.Join(wd, "temp", "example"))
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
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
		os.Stdout.WriteString(fmt.Sprintf("read request: %v", string(body)))
		(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("This is response")),
		}).Write(conn)
		conn.Close()
	}
}

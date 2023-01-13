package main

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	wd, _ := os.Getwd()
	address := path.Join(wd, "temp", "example")
	conn, err := net.Dial("unix", address)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	req, err := http.NewRequest("get", "http://localhost:80", strings.NewReader("body✋"))
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	req.Write(conn)
	time.Sleep(10 * time.Second)
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	dumped, err := httputil.DumpResponse(res, true)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	os.Stdout.WriteString(string(dumped))
}

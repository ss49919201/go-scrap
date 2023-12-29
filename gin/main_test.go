package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func Test_handleGetPing(t *testing.T) {
	r := gin.Default()
	handleGetPing(r)
	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	go func() {
		srv.ListenAndServe()
	}()
	defer srv.Close()

	time.Sleep(1 * time.Second)

	buf := bytes.NewBuffer(nil)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/ping", buf)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)

	require.JSONEq(t, `{"c.Get(\"ping\")":"pong","c.Request.Context().Value(\"ping\")":null,"message":"pong"}`, string(b))
}

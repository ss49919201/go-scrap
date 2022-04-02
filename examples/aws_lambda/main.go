package main

import (
	"github.com/go-tips/examples/aws_lambda/client"
	"github.com/go-tips/examples/aws_lambda/handler"
)

func main() {
	go handler.Run()
	client.Call()
}

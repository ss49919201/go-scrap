package main

import (
	"github.com/ss49919201/go-scrap/examples/aws_lambda/client"
	"github.com/ss49919201/go-scrap/examples/aws_lambda/handler"
)

func main() {
	go handler.Run()
	client.Call()
}

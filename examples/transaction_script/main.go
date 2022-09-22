package main

import (
	"fmt"

	"github.com/go-tips/examples/transaction_script/command"
	"github.com/go-tips/examples/transaction_script/model"
	"github.com/go-tips/examples/transaction_script/normal"
)

func main() {
	fmt.Println(normal.NewUserTS().Create("bar"))
	fmt.Println(normal.NewUserTS().FindByID(1))

	var runner command.Runner[*model.User]
	runner = command.NewCreateUserTS("bar")
	fmt.Println(runner.Run())
	runner = command.NewFindUserByID(1)
	fmt.Println(runner.Run())
}

package main

import (
	"fmt"

	"github.com/ss49919201/go-scrap/examples/transaction_script/command"
	"github.com/ss49919201/go-scrap/examples/transaction_script/model"
	"github.com/ss49919201/go-scrap/examples/transaction_script/normal"
)

func main() {
	// normal
	fmt.Println(normal.NewUserTS().Create("bar"))
	fmt.Println(normal.NewUserTS().FindByID(1))

	// command
	var runner command.Runner[*model.User]
	runner = command.NewCreateUserTS("bar")
	fmt.Println(runner.Run())
	runner = command.NewFindUserByID(1)
	fmt.Println(runner.Run())
}

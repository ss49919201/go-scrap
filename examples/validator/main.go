package main

import (
	"context"
	"fmt"

	_ "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
)

// 構造体はプライベートでもOK
type input struct {
	// フィールドはパブリックでなければならない
	Name string `validate:"required"`
	Age  string `validate:"required"`
}

func main() {
	err := validator.New().StructCtx(context.TODO(), input{})
	fmt.Println(err)
}

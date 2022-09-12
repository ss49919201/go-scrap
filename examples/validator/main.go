package main

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
)

func validate(ctx context.Context, v Validator, a any) error {
	return v.ValidateWithContext(ctx, a)
}

type Validator interface {
	ValidateWithContext(ctx context.Context, a any) error
}

// 構造体はプライベートでもOK
type input struct {
	// フィールドはパブリックでなければならない
	Name string `validate:"required"`
	Age  string `validate:"required"`
}

func (i *input) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, i,
		validation.Field(
			&i.Name,
			validation.Required,
		),
		validation.Field(
			&i.Age,
			validation.Required,
		),
	)
}

type validatorA struct{ *validator.Validate }

func (v validatorA) ValidateWithContext(ctx context.Context, a any) error {
	return v.StructCtx(ctx, a)
}

type validatorB struct{}

func (v validatorB) ValidateWithContext(ctx context.Context, a any) error {
	validatable, ok := a.(validation.ValidatableWithContext)
	if !ok {
		return nil
	}
	return validatable.ValidateWithContext(ctx)
}

func main() {
	fmt.Println("validatorA: ", validate(context.Background(), validatorA{validator.New()}, &input{}))
	fmt.Println("validatorB: ", validate(context.Background(), validatorB{}, &input{}))
}

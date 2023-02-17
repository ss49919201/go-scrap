package main

import (
	"errors"
	"fmt"

	"github.com/ss49919201/go-scrap/error/internal"
)

func main() {
	err1 := errors.New("error")
	err2 := &internal.CustomError{}
	err3 := fmt.Errorf("%w", &internal.CustomError{})

	fmt.Println(errors.Is(err1, err2))  // false
	fmt.Println(errors.As(err1, &err2)) // false

	fmt.Println(errors.Is(err2, err3))  // false
	fmt.Println(errors.As(err2, &err3)) // true

	fmt.Println(errors.Is(err3, err2))  // true
	fmt.Println(errors.As(err3, &err2)) // true

	err4 := internal.NewCustomError("custom")
	err5 := internal.WrapCustomError(errors.New("wrapped"))
	for _, err := range []error{err4, err5} {
		var ce *internal.CustomError
		if errors.As(err, &ce) {
			fmt.Printf("error = %s, stack = %s\n", ce.Error(), ce.Stack())
		}
	}
}

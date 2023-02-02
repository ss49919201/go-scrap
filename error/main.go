package main

import (
	"errors"
	"fmt"
)

type customError struct{}

func (c *customError) Error() string {
	return "customError"
}

func main() {
	err1 := errors.New("error")
	err2 := &customError{}
	err3 := fmt.Errorf("%w", &customError{})

	fmt.Println(errors.Is(err1, err2))  // false
	fmt.Println(errors.As(err1, &err2)) // false

	fmt.Println(errors.Is(err2, err3))  // false
	fmt.Println(errors.As(err2, &err3)) // true

	fmt.Println(errors.Is(err3, err2))  // true
	fmt.Println(errors.As(err3, &err2)) // true
}

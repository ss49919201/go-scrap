package internal

import (
	"fmt"
	"runtime"
)

type CustomError struct {
	msg   string
	stack string
}

func (c *CustomError) Error() string {
	return c.msg
}

func (c *CustomError) Stack() string {
	return c.stack
}

func NewCustomError(s string) error {
	return newCustomError(s, 1)
}

func WrapCustomError(err error) error {
	return newCustomError(err.Error(), 1)
}

func newCustomError(s string, skip int) *CustomError {
	pc, f, l, _ := runtime.Caller(skip + 1)
	fn := runtime.FuncForPC(pc)
	stack := fmt.Sprintf("%s %s: %d", fn.Name(), f, l)
	return &CustomError{
		msg:   s,
		stack: stack,
	}
}

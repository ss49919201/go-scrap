package internal

import (
	"fmt"
	"runtime"
)

type CustomError interface {
	error
	Stack() string
}

type customError struct {
	msg   string
	stack string
}

func (c *customError) Error() string {
	return c.msg
}

func (c *customError) Stack() string {
	return c.stack
}

func New(s string) error {
	return newCustomError(s, 1)
}

func WithStack(err error) error {
	return newCustomError(err.Error(), 1)
}

func newCustomError(s string, skip int) CustomError {
	pc, f, l, _ := runtime.Caller(skip + 1)
	fn := runtime.FuncForPC(pc)
	stack := fmt.Sprintf("%s %s: %d", fn.Name(), f, l)
	return &customError{
		msg:   s,
		stack: stack,
	}
}

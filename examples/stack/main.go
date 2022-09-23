package pattern

import (
	"errors"
	"reflect"
)

type Stack[T any] struct {
	sp   int
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		sp:   0,
		data: []T{},
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.sp == 0 {
		return zeroValue[T](), errors.New("stack is empty")
	}
	s.sp--
	return s.data[s.sp], nil
}

func (s *Stack[T]) Push(v T) error {
	if isNil(v) {
		return errors.New("value is invalid")
	}
	s.data = append(s.data, v)
	s.sp++
	return nil
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

func isNil(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
}

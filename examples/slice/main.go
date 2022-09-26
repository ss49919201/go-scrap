package main

import "fmt"

func popBack[T any](l []T) ([]T, T) {
	e := l[len(l)-1]
	l = l[:len(l)-1]
	return l, e
}

func ptr[T any](v T) *T {
	return &v
}

func sliceValue[T any](v []*T) []T {
	s := make([]T, len(v))
	for i, e := range v {
		s[i] = *e
	}
	return s
}

func main() {
	_3 := ptr(3)
	l := []*int{ptr(1), ptr(2), _3}
	var e *int

	l, e = popBack(l)
	fmt.Println(sliceValue(l), *e)
	*_3 = 5
	fmt.Println(sliceValue(l), *e)

	l, e = popBack(l)
	fmt.Println(sliceValue(l), *e)
	l, e = popBack(l)
	fmt.Println(sliceValue(l), *e)
}

package main

type Monad[T any] interface {
	Unit(v any) T
}

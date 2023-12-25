package main

import (
	"fmt"
	"iter"
)

func rangeOverFnA(yield func(i int, v string) bool) {
	b := yield(0, "0")
	fmt.Println(b) // true
}

func a() {
	for i, v := range rangeOverFnA {
		fmt.Println(i, v)
	}
}

func rangeOverFnB(yield func(i int, v string) bool) {
	b := yield(0, "0")
	fmt.Println(b) // false
}

func b() {
	for i, v := range rangeOverFnB {
		fmt.Println(i, v)
		break
	}
}

func rangeOverFnC(yield func(i int, v string) bool) {
	b := yield(0, "0")
	fmt.Println(b) // true
	if !b {
		return
	}
	b = yield(1, "1")
	fmt.Println(b) // false
	if !b {
		return
	}
	// not exec
	b = yield(2, "2")
	fmt.Println(b)
	if !b {
		return
	}
}

func c() {
	for i, v := range rangeOverFnC {
		fmt.Println(i, v)
		if i == 1 && v == "1" {
			break
		}
	}
}

func example[T any](seq iter.Seq[T]) {
	for i := range seq {
		fmt.Println(i)
	}
}

func example2[K, V any](seq iter.Seq2[K, V]) {
	for k, v := range seq {
		fmt.Println(k, v)
	}
}

func main() {
	example(func(yield func(int) bool) {
		if !yield(1) {
			return
		}
		if !yield(2) {
			return
		}
	})

	example2(func(yield func(int, string) bool) {
		if !yield(1, "1") {
			return
		}
		if !yield(2, "2") {
			return
		}
	})
}

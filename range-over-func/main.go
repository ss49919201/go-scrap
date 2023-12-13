package main

import "fmt"

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

func main() {}

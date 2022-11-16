package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var a string
	var b *string
	fmt.Println(reflect.DeepEqual(a, b))

	type T struct {
		A string
		B *string
		C []int
		D time.Time
	}
	b = new(string)
	c := T{
		A: "a",
		B: &a,
	}
	d := T{
		A: "a",
		B: b,
	}
	fmt.Println(reflect.DeepEqual(c, d))
}

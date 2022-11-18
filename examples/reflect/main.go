package main

import (
	"fmt"
	"reflect"
	"time"
)

func SameValueFiled[T any](x, y T) T {
	var res T
	valRes := reflect.ValueOf(&res)
	valX := reflect.ValueOf(x)
	valY := reflect.ValueOf(y)
	for i := 0; i < valX.NumField(); i++ {
		if reflect.DeepEqual(valX.Field(i).Interface(), valY.Field(i).Interface()) {
			valRes.Elem().Field(i).Set(valX.Field(i))
		}
	}
	return res
}

func DifficultValueFiled[T any](x, y T) T {
	var res T
	valRes := reflect.ValueOf(&res)
	valX := reflect.ValueOf(x)
	valY := reflect.ValueOf(y)
	for i := 0; i < valX.NumField(); i++ {
		if !reflect.DeepEqual(valX.Field(i).Interface(), valY.Field(i).Interface()) {
			valRes.Elem().Field(i).Set(valX.Field(i))
		}
	}
	return res
}

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
		A: "d",
		B: &a,
	}
	d := T{
		A: "a",
		B: b,
	}
	fmt.Println(reflect.DeepEqual(c, d))
	fmt.Println(SameValueFiled(c, d))
	fmt.Println(DifficultValueFiled(c, d))
}

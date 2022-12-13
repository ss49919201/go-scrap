package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
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

func toStringHeader(s string) *reflect.StringHeader {
	return (*reflect.StringHeader)(unsafe.Pointer(&s))
}

func main() {
	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int32(1)))
	m := make(map[int]int)
	m[1] = 1
	m2 := m
	m2[2] = 2
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m2)

	s := "string"
	fmt.Println(toStringHeader(s))
	s = ""
	fmt.Println(toStringHeader(s))
	s = "not string"
	fmt.Println(toStringHeader(s))

	bs := []byte("string")
	fmt.Println(toStringHeader(string(bs)))
	bs = bs[:0]
	fmt.Println(toStringHeader(string(bs)))
	bs = append(bs, []byte("s")...)
	fmt.Println(toStringHeader(string(bs)))

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

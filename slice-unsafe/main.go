package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type t struct {
	s []int
}

func (this t) imut() {
	p := unsafe.Pointer(&this.s)
	s := (*reflect.SliceHeader)(p)
	fmt.Println(s.Data)
	this.s[0] = 0
}

func (this *t) mut() {
	p := unsafe.Pointer(&this.s)
	s := (*reflect.SliceHeader)(p)
	fmt.Println(s.Data)
	this.s[0] = 0
}

func main() {
	tmp := t{}
	tmp.s = []int{1}

	p := unsafe.Pointer(&tmp.s)
	s := (*reflect.SliceHeader)(p)
	fmt.Println(s.Data)

	tmp.mut()
	fmt.Println(tmp.s)

	tmp.imut()
	fmt.Println(tmp.s)
}

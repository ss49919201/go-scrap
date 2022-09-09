package main

import (
	"reflect"
)

type parent struct{}

func (p *parent) foo() {
	println("parent foo")
}

type child struct {
	*parent
}

func (c *child) foo() {
	println("child foo")
}

func (c *child) bar() {
	println("child bar")
}

func (c *child) Foobar() {
	c.foo()
	c.bar()
}

func main() {
	c := child{}
	reflect.ValueOf(&c).MethodByName("Foobar").Call(nil) // child foo\nchild bar
	reflect.ValueOf(&c).Method(0).Call(nil)              // child foo\nchild bar
}

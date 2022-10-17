package main

import (
	"fmt"
)

// アルゴリズムを実行時に選択する

type Strategy func() bool

type Context struct {
	strategy Strategy
}

func (c *Context) Do() {
	fmt.Println(c.strategy())
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

type a struct{}

func (a a) Fn() bool {
	return false
}

func main() {
	c := NewContext(func() bool { return true })
	c.Do()

	c = NewContext(a{}.Fn)
	c.Do()
}

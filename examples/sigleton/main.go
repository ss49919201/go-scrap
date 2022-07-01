package main

import (
	"fmt"
	"sync"
)

var (
	instance    *Singleton
	instanceOne sync.Once
)

type Singleton struct {
	Name string
}

func New() *Singleton {
	instanceOne.Do(func() {
		instance = &Singleton{
			Name: "FirstValue",
		}
	})
	return instance
}

func main() {
	s := New()
	fmt.Println(s.Name) // FirstValue
	s.Name = "SecondValue"

	s1 := New()
	fmt.Println(s1.Name) // SecondValue
}

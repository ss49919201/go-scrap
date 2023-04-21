package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	info(print1, print2, print3, panic)
}

func info(fns ...any) {
	for _, fn := range fns {
		val := reflect.ValueOf(fn)
		name := runtime.FuncForPC(val.Pointer()).Name()
		fmt.Printf("😃 start %v\n", name)
		val.Call(nil)
		fmt.Printf("😃 end %v\n", name)
	}
}

func print1() {
	var m map[string][]int
	if m == nil {
		fmt.Println("`var m map[string][]int` is nil")
	}
	val := m["a"]
	fmt.Println(val)
}

func print2() {
	m := make(map[string][]int)
	val := m["a"]
	fmt.Println(val)
}

func print3() {
	m := make(map[string][]int, 5)
	m["a"] = append(m["a"], 1)
	fmt.Println(m)
}

func panic() {
	var m map[string][]int
	if m == nil {
		fmt.Println("`var m map[string][]int` is nil")
	}
	m["a"] = []int{}
	fmt.Println(m)
}

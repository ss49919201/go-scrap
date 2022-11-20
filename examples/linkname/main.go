package main

import (
	"fmt"
	_ "unsafe"
)

// go tool compile main.go
// go tool nm main.o
// go tool link -o a.out main.o

//go:linkname runtimeNano runtime.nanotime
func runtimeNano() int64

//go:linkname lower strconv.lower
func lower(c byte) byte

func main() {
	fmt.Println((runtimeNano()))
	fmt.Println(string(lower('A')))
	// Output:
	// 333328061583624
	// a
}

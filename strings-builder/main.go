package main

import (
	"fmt"
	"strings"
)

var builder strings.Builder

func main() {
	build()
	buildWithGrow()
}

func build() {
	defer builder.Reset()

	fmt.Printf("capacity=%d\n", builder.Cap())
	for i := 0; i < 1e9; i++ {
		builder.WriteByte(1)
	}
	fmt.Printf("length=%d\n", builder.Len())
}

func buildWithGrow() {
	defer builder.Reset()

	builder.Grow(1e9)
	fmt.Printf("capacity=%d\n", builder.Cap())
	for i := 0; i < 1e9; i++ {
		builder.WriteByte(1)
	}
	fmt.Printf("length=%d\n", builder.Len())
}

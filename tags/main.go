package main

import (
	"github.com/ss49919201/go-scrap/tags/internal"
)

func tags() string {
	return internal.Tags()
}

func main() {
	println(internal.Tags())
}

package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	fmt.Printf("%+v", xerrors.New("hello"))
	// hello:
	// main.main
	//     /Users/sakaeshinya/workspace/go-scrap/examples/xerrors/main.go:13%
}

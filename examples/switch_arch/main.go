package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOARCH {
	case `arm64`:
		fmt.Println("arm64")
	case `amd64`:
		fmt.Println("amd64")
	}
}

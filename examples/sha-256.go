package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("こんにちは世界"))
	fmt.Printf("%x", h.Sum(nil))
}

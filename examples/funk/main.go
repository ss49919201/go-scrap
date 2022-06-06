package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

type hoge struct{}

func main() {
	hl := []*hoge{{}, {}, {}}
	v := funk.Find(hl, func(h *hoge) bool {
		return false
	}).(*hoge)
	fmt.Println(v == nil)
}

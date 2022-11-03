package main

import (
	"fmt"

	"github.com/ahmetb/go-linq/v3"
)

func main() {
	var s []string
	linq.From([]string{"1", "2"}).ToSlice(&s)
	fmt.Println(s)
}

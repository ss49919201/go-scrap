package main

import (
	"fmt"

	"github.com/ss49919201/go-scrap/examples/linq_generic/internal"
)

func main() {
	q := internal.From([]string{"1", "2"})
	fmt.Println(q.ToSlice())

	fmt.Println(
		q.Where(func(s string) bool {
			return s == "1"
		}).ToSlice(),
	)
}

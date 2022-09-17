package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	reduce()
}

func reduce() {
	fmt.Println(
		lo.Reduce(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			func(acc string, x, _ int) string {
				acc += strconv.Itoa(x)
				return acc
			},
			"This is the accumulator: ",
		),
	)
}

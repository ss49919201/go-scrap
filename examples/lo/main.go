package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	uselo()
}

func useStandard() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func uselo() {
	fmt.Println(
		lo.Reduce([]int{1, 2, 3, 4, 5}, func(acc string, x, _ int) string {
			acc += strconv.Itoa(x)
			return acc
		}, ""),
	)
}

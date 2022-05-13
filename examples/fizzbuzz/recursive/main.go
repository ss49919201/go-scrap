package main

import (
	"fmt"
	"strconv"
)

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	do(1, 50)
}

func do(i, max int) {
	var result string
	if i > max {
		return
	}
	if isFizz(i) {
		result += fizz
	}
	if isBuzz(i) {
		result += buzz
	}
	if len(result) == 0 {
		result = strconv.Itoa(i)
	}
	fmt.Println(result)

	do(i+1, max)
}

func isFizz(i int) bool {
	return i%3 == 0
}

func isBuzz(i int) bool {
	return i%5 == 0
}

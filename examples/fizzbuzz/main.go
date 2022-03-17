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
	for i := 1; i <= 50; i++ {
		var result string
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
	}
}

func isFizz(i int) bool {
	return i%3 == 0
}

func isBuzz(i int) bool {
	return i%5 == 0
}

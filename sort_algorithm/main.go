package main

import (
	"fmt"
)

var target = []int{3, 1, 4, 5, 2}

func main() {
	fmt.Println(bubble(target))
	fmt.Println(merge(target))
}

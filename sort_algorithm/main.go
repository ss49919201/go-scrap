package main

import (
	"fmt"
)

var target = []int{3, 1, 4, 5, 2}

func main() {
	fmt.Println(bubble(target))
}

// バブルソート
// 最悪 O(N^2)
func bubble(target []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := range target {
			if i == len(target)-1 {
				break
			}

			if target[i] > target[i+1] {
				target[i], target[i+1] = target[i+1], target[i]
				swapped = true
			}
		}
	}
	return target
}

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	a := []int{3, 2, 3, 0, 1, 3, 5, 9, 4, 10, 7, 8, 6}
	b := make([]int, len(a))
	copy(b, a)
	quicksort(a)
	quickSort(b, 0, len(b)-1)
	fmt.Println(a)
	fmt.Println(b)
}

func quickSort[T constraints.Ordered](arr []T, lo, hi int) {
	partition := func(lo, hi int) int {
		pivot := arr[hi]
		pivotIndex := lo
		for j := lo; j < hi; j++ {
			if arr[j] < pivot {
				arr[pivotIndex], arr[j] = arr[j], arr[pivotIndex]
				pivotIndex++
			}
		}
		arr[pivotIndex], arr[hi] = arr[hi], arr[pivotIndex]
		return pivotIndex
	}

	if lo < hi {
		p := partition(lo, hi)
		quickSort(arr, lo, p-1)
		quickSort(arr, p+1, hi)
	}
}

func quicksort[T constraints.Ordered](arr []T) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[len(arr)-1]
	pivotIndex := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[pivotIndex], arr[i] = arr[i], arr[pivotIndex]
			pivotIndex++
		}
	}
	arr[pivotIndex], arr[len(arr)-1] = pivot, arr[pivotIndex]
	quicksort(arr[:pivotIndex])
	quicksort(arr[pivotIndex+1:])
}

package main

import "sort"

func merge(target []int) []int {
	half := len(target) / 2
	if half == 0 {
		return target
	}

	a := merge(target[:half])
	b := merge(target[half:])

	tmpSlice := make([]int, 0, len(target))
	tmpSlice = append(tmpSlice, a...)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	tmpSlice = append(tmpSlice, b...)

	begin := 0
	end := len(tmpSlice) - 1
	for i := 0; i < len(tmpSlice); i++ {
		if tmpSlice[begin] < tmpSlice[end] {
			target[i] = tmpSlice[begin]
			begin++
		} else {
			target[i] = tmpSlice[end]
			end--
		}
	}
	return target
}

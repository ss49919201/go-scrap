package main

// クイックソート
func quick(target []int) []int {
	if len(target) < 2 {
		return target
	}
	pivot := target[len(target) - 1]
	var l, r []int
	for i := 0; i < len(target)-1; i++ {
		if target[i] <= pivot {
			l = append(l, target[i])
		} else {
			r = append(r, target[i])
		}
	}
	return append(append(quick(l), pivot), quick(r)...)
}

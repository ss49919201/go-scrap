package main

func linear(target []int, want int) int {
	size := len(target)
	target = append(target, want)
	i := 0
	for target[i] != want {
		i++
	}
	if i == size {
		return -1
	}
	return want
}

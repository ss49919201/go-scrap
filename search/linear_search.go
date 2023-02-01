package main

func linear(target []int, want int) int {
	for i := 0; i < len(target); i++ {
		if target[i] == want {
			return target[i]
		}
	}
	return -1
}

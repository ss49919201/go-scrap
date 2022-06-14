package main

func fibonach(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}

	return fibonach(n-2) + fibonach(n-1)
}

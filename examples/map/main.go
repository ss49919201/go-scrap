package main

import "fmt"

func main() {
	m := make(map[int][]int)
	pushBack(m, 1, 1)
	fmt.Println(m) // map[1:[1]]
	pushBack(m, 1, 2)
	fmt.Println(m) // map[1:[1 2]]
}

func pushBack(m map[int][]int, key int, value int) {
	m[key] = append(m[key], value)
}

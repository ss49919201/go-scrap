package main

import "fmt"

const (
	A uint = 10 // 1010
	B uint = 12 // 1100
)

func main() {
	var bits uint

	// AND演算
	bits = A & B             // 1001
	fmt.Printf("%b\n", bits) // 1001
	// OR演算
	// bits = A | B // 1110
	// // XOR演算
	// bits = A ^ B // 0110
	// // AND NOT演算
	// bits = A &^ B // 0010
	// // 左シフト演算
	// bits = 1 << uint64(3) // 1000 : 2の3乗かかる
	// // 右シフト演算
	// bits = 8 >> uint64(3) // 0001 : 2の(-3)乗かかる
}

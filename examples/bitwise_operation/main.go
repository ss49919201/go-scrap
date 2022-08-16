package main

import "fmt"

func main() {
	fmt.Printf("%08b\n", 1<<uint(1)) // 00000010

	bit := 5
	for i := 0; i < (1 << bit); i++ {
		fmt.Printf("%08b\n", i)
		fmt.Printf("%08b\n", 1<<bit)
	}
}

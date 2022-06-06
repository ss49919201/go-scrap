package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo()
	println(concat([]string{"hoge", "fuga"}))
}

func echo() {
	fmt.Println("os.Args")
	fmt.Println(strings.Join(os.Args, " "))
}

func concat(stringSlice []string) string {
	s, sep := "", ""
	for _, str := range stringSlice {
		s += sep + str
		sep = string(rune(0xA0))
	}
	return s
}

// gpl/ch01/ex04/dup2.go からやる

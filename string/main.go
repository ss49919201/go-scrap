package main

import (
	"fmt"
	"strings"
)

func main() {
	sb := strings.Builder{}
	for _, v := range toRuneSlice("あいうえお💫") {
		sb.WriteString(string(v))
	}
	fmt.Println(sb.String()) // あいうえお💫

	fmt.Println(toString([]rune{1000, 2000, 3000})) // Ϩߐஸ
}

func toRuneSlice(s string) []rune {
	result := []rune{}
	for _, v := range s {
		result = append(result, v)
	}
	return result
}

func toString(rs []rune) string {
	return string(rs)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanText() string {
	sc.Scan()
	return sc.Text()
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func toIntList(s []string) []int {
	ns := make([]int, len(s))
	for i, v := range s {
		ns[i] = toInt(v)
	}
	return ns
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeDash(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	N := scanText()
	ss := strings.Split(N, " ")
	h1 := toInt(ss[0])
	_ = toInt(ss[1])
	as := make([][]int, h1)
	for i := 0; i < h1; i++ {
		N := scanText()
		ss := strings.Split(N, " ")
		as[i] = toIntList(ss)
	}

	N2 := scanText()
	ss2 := strings.Split(N2, " ")
	h2 := toInt(ss2[0])
	_ = toInt(ss2[1])
	bs := make([][]int, h2)
	for i := 0; i < h2; i++ {
		N := scanText()
		ss := strings.Split(N, " ")
		bs[i] = toIntList(ss)
	}
	fmt.Println(as, bs)
}

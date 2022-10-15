package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const BUFSIZE = 1000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type mem map[int]int

func main() {
	N := toInt(readline())
	A := readIntSlice()
	m := make(mem)
	m2 := make(mem)

	tmpM := make(mem)
	uA := make([]int, 0)
	for _, v := range A {
		if _, ok := tmpM[v]; !ok {
			uA = append(uA, v)
			tmpM[v] = 1
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(uA)))

	for i := 0; i < len(uA); i++ {
		m[uA[i]] = i
	}

	for i := 0; i < N; i++ {
		m2[m[A[i]]]++
	}

	for i := 1; i <= N; i++ {
		fmt.Println(m2[i-1])
	}
}

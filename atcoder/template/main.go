package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func mod(n1, n2 int) int {
	res := (n1 + n2) % n2
	if res < 0 {
		res += n2
	}
	return res
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func popBack[T any](l *[]T) T {
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e
}

type mem map[int]int

func main() {
	N := toInt(readline())
	NS := readIntSlice()
	m := new(mem)

	fmt.Println(N, NS, m)
}

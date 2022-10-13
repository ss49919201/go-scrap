//go:build ignore

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

var mem = make(map[int]int)

func main() {
	NK := readIntSlice()
	N := NK[0]
	K := NK[1]
	NS := readIntSlice()

	mem[1] = 0
	mem[2] = int(math.Abs(float64(NS[1]) - float64(NS[0])))

	for i := 2; i < N; i++ {
		j := i - K
		if j < 0 {
			j = 0
		}

		min := -1
		for ; j < i; j++ {
			t := int(math.Abs(float64(NS[i])-float64(NS[j])) + float64(mem[j+1]))
			if min == -1 {
				min = t
				continue
			}
			if t < min {
				min = t
			}
		}
		mem[i+1] = min
	}

	fmt.Println(mem[N])
}

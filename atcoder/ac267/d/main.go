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
	N := toInt(readline())
	NS := readIntSlice()

	mem[1] = 0
	mem[2] = int(math.Abs(float64(NS[1]) - float64(NS[0])))

	for i := 2; i < N; i++ {
		// 1つずつ
		a := math.Abs(float64(NS[i])-float64(NS[i-1])) + float64(mem[i])
		// 1つとばし
		b := math.Abs(float64(NS[i])-float64(NS[i-2])) + float64(mem[i-1])

		mem[i+1] = int(math.Min(a, b))
	}

	fmt.Println(mem[N])
}

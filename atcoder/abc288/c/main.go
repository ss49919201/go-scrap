package main

import (
	"bufio"
	"fmt"
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

func main() {
	N := toInt(readline())
	PS := readIntSlice()

	// 回数ごとの最大値
	count := make([]int, N)

	// 全部の料理の喜ばれる回数を探す
	for i, v := range PS {
		for j := 0; j < 3; j++ {
			count[(v-i+N-1+j)%N]++
		}
	}

	var res int
	for _, v := range count {
		if v > res {
			res = v
		}
	}

	fmt.Println(res)
}

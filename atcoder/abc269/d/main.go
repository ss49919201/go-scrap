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

type mem map[vector]bool

type vector struct {
	x, y int
}

var m = make(mem)
var m2 = make(mem)

func main() {
	N := toInt(readline())
	XY := make([]vector, N)
	for i := 0; i < N; i++ {
		xy := readIntSlice()
		XY[i] = vector{xy[0], xy[1]}
	}

	ans := 0
	check := func(v vector) {
		m[vector{v.x - 1, v.y - 1}] = true
		m[vector{v.x - 1, v.y}] = true
		m[vector{v.x, v.y - 1}] = true
		m[vector{v.x, v.y + 1}] = true
		m[vector{v.x + 1, v.y}] = true
		m[vector{v.x + 1, v.y + 1}] = true
	}
	for _, v := range XY {
		if ok := m[v]; !ok {
			m[v] = true
			ans++
		}

		f(v, XY, check)
	}

	fmt.Println(ans)
}

func f(v vector, xy []vector, check func(v vector)) {
	check(v)
	for _, v2 := range xy {
		if v == v2 {
			m2[v] = true
			continue
		}
		if ok := m[v2]; ok {
			if ok := m2[v2]; !ok {
				m2[v2] = true
				f(v2, xy, check)
			}
		}
	}
}

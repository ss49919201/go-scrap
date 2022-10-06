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

type mem map[int]int

var flag = make(map[int]bool)
var e = make(map[int][]int)
var deque []int
var stop bool

func dfs(cur, next int) {
	// duque の末尾に cur を追加
	if stop {
		return
	} else {
		deque = append(deque, cur)
	}

	// 目的地に到達したら終了
	if cur == next {
		stop = true
		return
	}

	// 頂点到達フラグを更新
	flag[cur] = true

	// 現在地から行ける頂点を探索
	// 既に到達済みの頂点は探索しない
	for _, v := range e[cur] {
		if !flag[v] {
			dfs(v, next)
		}
	}

	// 現在地から行ける頂点全てが訪問済み又は行き止まりの場合は
	// 末尾の頂点を削除 (つまり、一つ前の頂点での探索に戻る)
	if stop {
		return
	} else {
		popBack(&deque)
	}
}

func popBack(l *[]int) int {
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e
}

func main() {
	l1 := readIntSlice()
	N, X, Y := l1[0], l1[1], l1[2]

	for i := 0; i < N-1; i++ {
		l := readIntSlice()
		e[l[0]] = append(e[l[0]], l[1])
		e[l[1]] = append(e[l[1]], l[0])
	}

	dfs(X, Y)

	for i, v := range deque {
		fmt.Print(v)
		if (i + 1) != len(deque) {
			fmt.Print(" ")
		}
	}
}

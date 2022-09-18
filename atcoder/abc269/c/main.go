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

func mod(n1, n2 int) int {
	res := (n1 + n2) % n2
	if res < 0 {
		res += n2
	}
	return res
}

type mem map[int]int

func main() {
	// N=11
	N := toInt(readline())

	// 1が立っている位を調べる
	// 最小の位は0
	// {0,1,3}
	A := make([]int, 0)
	for i := 0; i <= 60; i++ {
		if (1<<i)&N > 0 {
			A = append(A, i)
		}
	}

	// {0,1,3}の全ての組み合わせが答え
	// ２の3乗で8通りなので全探索
	// {0,1,3}->{001,010,100}として考えるので、{1<<0,1<<1,1<<2}
	// 0~7(000~111)それぞれについて、1が立っている位の和を計算する
	// 例: i=1->1=001->001&1<<0=1,001&1<<1=0,001&1<<2=0 -> {0,1,3}のうち{0}だけたっているので、2の0乗=1を答えに追加
	// 例: i=4->4=100->100&1<<0=0,100&1<<1=0,100&1<<2=1 -> {0,1,3}のうち{3}だけたっているので、2の3乗=8を答えに追加
	// 例: i=6->4=110->110&1<<0=0,110&1<<1=1,110&1<<2=1 -> {0,1,3}のうち{1,3}だけたっているので、2の3乗+2の1乗=10を答えに追加
	K := len(A)
	Ans := make([]int, 0)
	for i := 0; i < (1 << K); i++ {
		cur := 0
		for j := 0; j < K; j++ {
			if (1<<j)&i > 0 {
				cur |= 1 << A[j]
			}
		}
		Ans = append(Ans, cur)
	}

	sort.Ints(Ans)
	for _, v := range Ans {
		fmt.Println(v)
	}
}

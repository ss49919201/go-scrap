package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 最適な組み合わせを考える際に、各段階での最適解を選択し、それらを組み合わせて最終的な解とする。

const BUFSIZE = 1000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

// https://algo-method.com/tasks/359
func main() {
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
	NS := string(buf)
	N, _ := strconv.Atoi(NS)

	var ans int
	// Step1: 5円玉を最大使用可能枚数
	ans += N / 5
	// Step2: 1円玉の最大使用可能枚数
	ans += N % 5
	// Step1,Step2の組み合わせが最適解になる
	fmt.Println(ans)
}

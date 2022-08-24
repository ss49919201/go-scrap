//go:build ignore

package main

import (
	"sort"
)

// 同じ探索を複数回しない

// 容量
var U = 10

// 品物の数
var N = 10

// 品物の重さ
var W = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 品物の価値
var V = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func solveNoDP(i, u int) int {
	if i == N {
		// これ以上品物はない
		return 0
	}

	// この品物は容量を超える為、次の品物へ
	if W[i] > u {
		return solveNoDP(i+1, u)
	}

	// スキップした場合と、選んだ場合の価値を比較
	s := sort.IntSlice([]int{
		solveNoDP(i+1, u),             // スキップした場合
		solveNoDP(i+1, u-W[i]) + V[i], // 選んだ場合
	})
	sort.Sort(s)
	return s[1]
}

var mem = make(map[struct{ i, u int }]int)

func solveWithDP(i, u int) int {
	if i == N {
		// これ以上品物はない
		return 0
	}

	if v, ok := mem[struct {
		i int
		u int
	}{i, u}]; ok {
		return v
	}

	// この品物は容量を超える為、次の品物へ
	if W[i] > u {
		r := solveNoDP(i+1, u)

		mem[struct {
			i int
			u int
		}{i, u}] = r

		return r
	}

	// スキップした場合と、選んだ場合の価値を比較
	s := sort.IntSlice([]int{
		solveNoDP(i+1, u),             // スキップした場合
		solveNoDP(i+1, u-W[i]) + V[i], // 選んだ場合
	})
	sort.Sort(s)
	r := s[1]

	mem[struct {
		i int
		u int
	}{i, u}] = r

	return r
}

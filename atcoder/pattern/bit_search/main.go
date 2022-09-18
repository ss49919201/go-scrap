package main

import "fmt"

func main() {
	// N個の要素から0~N個選ぶパターンを列挙して調べる
	// 例えばN=3のとき、2^3=8パターン

	N := [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	W := 6
	Ans := [][]int{}

	// 2^20=1048576通りのbit列を全探索
	// 0~1048575
	for i := 0; i < 1<<len(N); i++ {
		sum := 0
		pattern := []int{}

		for j := 0; j < len(N); j++ {
			// jbit目が1であれば、j番目の要素を加算
			if 1<<j&i > 0 {
				sum += N[j]
				pattern = append(pattern, N[j])
			}
		}

		if sum == W {
			Ans = append(Ans, pattern)
		}
	}

	for _, v := range Ans {
		fmt.Println(v)
	}

	// Output:
	// [1 2 3]
	// [0 1 2 3]
	// [2 4]
	// [0 2 4]
	// [1 5]
	// [0 1 5]
	// [6]
	// [0 6]
}

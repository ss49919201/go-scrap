package main

import "fmt"

func dx() [8]int {
	return [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
}
func dy() [8]int {
	return [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
}

func solve(n int, rows [][]int64) int64 {
	var ans int64

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			// 8方向を全探索
			for d := 0; d < 8; d++ {
				var tmp int64
				for i := 0; i < n; i++ {
					tmp *= 10
					var rowsY, rowsX int

					// y は初期値
					// dy[d]は移動方向
					// i は移動回数
					// n で割った余りが移動先の値
					// n で割ることで最大値をはみ出た時に、0から余り分進んだ値になる
					rowsY = (y + dy()[d]*i) % n

					// 負数の場合は、最大値から絶対値を引いて求める
					// n = 4, rowsY = -1, rowsY = 3
					if rowsY < 0 {
						rowsY += n
					}

					// x は初期値
					// dy[d]は移動方向
					// i は移動回数
					// n で割った余りが移動先の値
					// n で割ることで最大値をはみ出た時に、0から余り分進めた値になる
					rowsX = (x + dx()[d]*i) % n
					if rowsX < 0 {
						rowsX += n
					}

					tmp += rows[rowsY][rowsX]
				}
				if ans < tmp {
					ans = tmp
				}
			}
		}
	}

	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	rows := make([][]int64, n)
	for i := 0; i < n; i++ {
		rows[i] = make([]int64, n)
		var a string
		fmt.Scan(&a)
		for j := 0; j < n; j++ {
			rows[i][j] = int64(a[j] - '0')
		}
	}

	fmt.Println(solve(n, rows))
}

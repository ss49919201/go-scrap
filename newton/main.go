package main

import (
	"fmt"
	"math"
)

func main() {
	solve()
}

func solve() {
	// ルート2を求めたいので、
	// y=x^2の式の、y=2の時のxを求めれば良い
	r := 2.0
	a := 2.0

	// y=x^2の式
	f := func(n float64) float64 { return math.Pow(n, 2) }
	// 導関数
	fDash := func(n float64) float64 { return 2 * n }

	for i := 0; i < 10; i++ {
		// 直線の方程式より接線の方程式を求める
		// y-b=m(x-a) -> x = (y - b + ma) / m
		// aを交点のx座標に上書きする
		a = (r - f(a) + (fDash(a) * a)) / fDash(a)
		fmt.Println(a)
	}
}

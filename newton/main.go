package main

import (
	"fmt"
)

func main() {
	solve()
}

func solve() {
	// ルート2を求めたいので、
	// y=x^2の式の、y=2の時のxを求めれば良い
	r := 2.0
	a := 2.0

	for i := 0; i < 10; i++ {
		zahyouX := a
		zahyouY := a * a

		// 接線の方程式を求める
		katamuki := 2.0 * zahyouX            // 微分の公式から傾きが求められる
		seppen := zahyouY - katamuki*zahyouX // 切片
		// 求めた接線の方程式のyにrを代入することで、y=rとの交点のx座標を求める
		a = (r - seppen) / katamuki
		fmt.Println(a)
	}
}

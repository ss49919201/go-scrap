package main

func main() {
	add1 := add(1)
	println(add1(1))
	println(add1(2))
}

// 複数の引数を取る関数を一つの引数を取る関数に変換するとを、カリー化という
// 第一引数が1の場合の関数を使いまわせる
// add(1, hoge), add(1, fuga) ... の add(1 の重複を防げる
func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

package pattern

import (
	"fmt"
	"os"
)

// return 文に比較演算子
func RetBool() bool {
	t := true
	f := false
	return t == f
}

func Switch(parm int) {
	switch parm {
	// defaultはどこに書いてもいい
	default:
		fmt.Println("default")
	case 1:
		fmt.Println("true")
	// OR条件で書ける
	case 2, 3:
		fmt.Println("false")
	}
}

func CreateFile(name string) {
	// ./ にファイルを作成
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f.Name())

	// 掃除
	defer os.Remove(f.Name())
}

func CreateTempFile(name string) {
	// 第1引数を空文字列にすることで 環境変数TMPDIR にファイルを作成
	f, err := os.CreateTemp("", name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f.Name())

	// 掃除
	defer os.Remove(f.Name())
}

func AppendSlice() {
	// FIXME: capを確保する
	l := []int{1, 2, 3, 4}
	r := []int{1, 2, 3, 4}
	l = append(l, r...)
	fmt.Println(l)
}

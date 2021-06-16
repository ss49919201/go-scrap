package pattern

import "fmt"

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

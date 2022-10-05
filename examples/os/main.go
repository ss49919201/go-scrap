package main

import (
	"fmt"
	"os"
)

func main() {
	stat()
}

// ファイルの情報を確認する
// ファイルの存在チェックなどにも使われる
func stat() {
	fmt.Println(os.Stat("examples"))
}

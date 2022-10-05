package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	walk()
}

func join() {
	fmt.Println(filepath.Join("a/b", "../../../xyz"))
}

// 指定のディレクトリ以下のファイルを再帰的に取得し、クロージャーを実行する
func walk() {
	filepath.Walk("./examples/filepath", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(path, info.IsDir())
		return nil
	})
	// ./examples/filepath true
	// examples/filepath/main.go false
}

func glob() {

}

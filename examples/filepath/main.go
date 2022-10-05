package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	join()
	walk()
	glob()
}

func join() {
	fmt.Println("---------- join ----------")
	fmt.Println(filepath.Join("a/b", "../../../xyz"))
}

// 指定のディレクトリ以下のファイルを再帰的に取得し、クロージャーを実行する
func walk() {
	filepath.Walk("./examples/filepath", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println("---------- walk ----------")
		fmt.Println(path, info.IsDir())
		return nil
	})
	// ./examples/filepath true
	// examples/filepath/main.go false
}

// 引数のパターンにマッチしたファイル名を取得する
func glob() {
	files, err := filepath.Glob("./examples/filepath/*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("---------- glob ----------")
	for _, f := range files {
		fmt.Println(f)
	}
	// examples/filepath/dummy.go
	// examples/filepath/main.go
}

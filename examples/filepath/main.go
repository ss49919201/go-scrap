package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	join()
	walk()
	glob()
	dir()
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

// ファイルを探してのディレクトリパスを取得する
func dirInternal(file string) (string, bool) {
	_, filename, _, _ := runtime.Caller(0)
	fn := filepath.Dir
	for dir := fn(filename); dir != "/" && dir != "."; dir = fn(dir) {
		_, err := os.Stat(filepath.Join(dir, file))
		if err == nil {
			return dir, true
		}
	}
	return "", false
}

func dir() {
	s, ok := dirInternal(".gitignore")
	if ok {
		fmt.Println("---------- dir ----------")
		fmt.Println(s)
	}
}

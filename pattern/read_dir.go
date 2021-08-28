package pattern

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: サブディレクトリの対応
func ReadDir(dir string) {
	// ディレクトリのオープン
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// ディレクトリ内のファイル情報を取得
	files, err := f.ReadDir(0)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		// 情報を出力
		fmt.Println(file.Name())
		fmt.Println(file.Type().IsRegular())
		fmt.Println(file.Type().IsDir())
		fmt.Println(file.Info())

		// ファイルパスを出力
		d := filepath.Join(dir, file.Name())
		fmt.Println(d)

		// ファイルをオープン
		f, err := os.Open(d)
		if err != nil {
			panic(err)
		}

		// テキストを出力
		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(s.Text())
		}

		fmt.Println()
	}
}

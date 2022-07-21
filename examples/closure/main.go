package main

import "fmt"

// > クロージャ（クロージャー、英語: closure）、関数閉包はプログラミング言語における関数オブジェクトの一種。
// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%B8%E3%83%A3

// 引数以外の変数を実行時の状態ではなく、
// 定義されたスコープで扱うことができる。

type store struct {
	k string
	v any
}

type set struct{}

func (s *set) exec(store store, k string, v any) {}

type del struct{}

func (d *del) exec(store store, k string, v any) {}

type cmd interface {
	exec(s store, k string, v any)
}

func storeCmd() func(c cmd, k string, v any) {
	var s store
	return func(c cmd, k string, v any) {
		c.exec(s, k, v)
		fmt.Println(s)
	}
}

func main() {
	cmd1 := storeCmd()
	cmd1(&set{}, "k", "v")
	fmt.Println()
}

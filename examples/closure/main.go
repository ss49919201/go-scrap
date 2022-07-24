package main

import "fmt"

// > クロージャ（クロージャー、英語: closure）、関数閉包はプログラミング言語における関数オブジェクトの一種。
// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%B8%E3%83%A3

// 引数以外の変数を実行時の状態ではなく、
// 定義されたスコープで扱うことができる。

type userType string

const (
	teacher userType = "teacher"
	student userType = "student"
)

type user struct {
	typ userType
}

func (u *user) greet() {
	if u.typ == teacher {
		fmt.Println("Hello teacher!")
	} else {
		fmt.Println("Hello student!")
	}
}

func (u *user) greetFunc() func() {
	var msg string
	if u.typ == teacher {
		msg = "Hello teacher!"
	} else {
		msg = "Hello student!"
	}
	return func() {
		fmt.Println(msg)
	}
}

func counterFunc() func() int {
	var n int
	return func() int {
		n++
		return n
	}
}

func main() {
	u := &user{teacher}
	// 毎回typの評価が実行される
	u.greet()
	u.greet()

	// typの評価は一度だけ実行される
	f := u.greetFunc()
	f()
	f()

	cf := counterFunc()
	fmt.Println(cf()) // 1
	fmt.Println(cf()) // 2

	cf2 := counterFunc()
	fmt.Println(cf2()) // 1
	fmt.Println(cf())  // 3
}

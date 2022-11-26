package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(true))
	// チャネルはgoroutine間で値をやり取りする為の仕組みである
	// make関数で作成される
	// データ構造への参照型である
	var c chan string
	println(c == nil) // zero値はnil
	c = make(chan string)

	// 引数で渡した場合も同じデータ構造を参照する
	go func(copiedC chan string) {
		// 同じデータ構造への参照であれば比較は真
		if c == copiedC {
			time.Sleep(1 * time.Second)
			copiedC <- "Hello, World!"
		}
	}(c)

	// channel から値を受け取るまでブロックされる
	println(<-c)

	// バッファなしのチャネルを使う場合、受信が始まるまで、送信はブロックされる
	// c <- "Blocked" // fatal error: all goroutines are asleep - deadlock!
	c = make(chan string, 1)
	c <- "Blocked"

	// これ以上送信がない時はcloseするのが良い
	func(copiedC chan string) {
		if c == copiedC {
			close(copiedC)
		}
	}(c)
	// closeした後に送信するとpanic
	// func(copiedC chan string) {
	// 	if c == copiedC {
	// 		copiedC <- "Close?" // panic: send on closed channe
	// 	}
	// }(c)
	// closeした後でも、データ構造から値が空になるまでは受信可能
	println(<-c)
}

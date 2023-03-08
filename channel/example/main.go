package main

import (
	"fmt"
	"time"
)

func main() {
	traceLifeCycle()
}

// @for range 句はchannelがcloseした時に終了すること
// MEMO: closeしなかった場合、for range句は読み込みブロックする
func traceForRange() {
	c := make(chan int, 4)
	for i := 0; i < 4; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}

// @バッファなしchannelへの書き込みから受信ができること
func traceLifeCycle() {
	c := make(chan string)
	go func() {
		// バッファなしchannelは初期化時点で満杯扱いなので
		// 受信されるまで書き込みがブロックされる
		// goroutineにすることで書き込みのブロックを回避する
		c <- "trace"
	}()
	v, ok := <-c
	fmt.Println(v, ok)
}

// @closeされていても読み込みは可能であること
// MEMO: channelが空になって初めて、ゼロ値とfalseが返る
func traceReadingClosedChannel() {
	c := make(chan string)
	go func() {
		// バッファなしchannelは初期化時点で満杯扱いなので
		// 受信されるまで書き込みがブロックされる
		// goroutineにすることで書き込みのブロックを回避する
		c <- "trace"
		close(c)
	}()

	time.Sleep(time.Second)
	_, ok := <-c
	fmt.Println(ok)
	_, ok = <-c
	fmt.Println(ok)
}

func fetchAndPrint() {
	data := asyncFetch()
	fmt.Println(<-data) // data取得までブロック
}

// send,closeをできないようにreceived-only channelを返す
func asyncFetch() <-chan string {
	stream := make(chan string)
	go func() {
		time.Sleep(time.Second)
		fetchAndWrite(stream)
		close(stream)
	}()
	return stream
}

// send-onlyからreceived-onlyへは変換不可なのでsend-onlyを返す
// MEMO: 受信先が既にあれば送信元から透過的に値が渡される
func fetchAndWrite(stream chan<- string) chan<- string {
	stream <- store[1]
	return stream
}

var store = map[int]string{
	1: "fetched data",
}

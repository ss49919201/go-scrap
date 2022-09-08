package main

import "time"

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c <- "Hello, World!"
	}()

	// channel から値を受け取るまでブロックされる
	println(<-c)
}

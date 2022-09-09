package main

import (
	"fmt"
	"sync"
	"time"
)

func badParallel() {
	var wg sync.WaitGroup
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello, World!")
		wg.Add(1)
	}()
	wg.Wait() // 待たずに終わる
}

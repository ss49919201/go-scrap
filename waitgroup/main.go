package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var wg sync.WaitGroup
	successChan := make(chan int, 5)
	type faild struct {
		userID int
		err    error
	}
	failedChan := make(chan faild, 5)
	for _, v := range []int{1, 2, 3, 4, 5} {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			err := updateUser(userID)
			if err == nil {
				successChan <- userID
			} else {
				failedChan <- faild{userID, err}
			}
		}(v)
	}
	wg.Wait()

	// Closeしておかないと for 文の range 句の右側の式に、channel を渡した際に channel への送信を待ち続けてデッドロックになる
	close(successChan)
	close(failedChan)

	for v := range successChan {
		fmt.Println("success: ", v)
	}

	for v := range failedChan {
		fmt.Println("failed: ", v.err)
	}

	return nil
}

func updateUser(userID int) error {
	if rand.Intn(5) == userID {
		return fmt.Errorf("invalid user id: %d", userID)
	}
	return nil
}

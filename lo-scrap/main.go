package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	doTransaction()
}

func doTransaction() {
	tx := lo.NewTransaction[int]().Then(
		func(n int) (int, error) {
			fmt.Println("Exec1", n)
			return n + 1, nil
		},
		func(n int) int {
			fmt.Println("Rollback1", n)
			return n + 1
		},
	).Then(
		func(n int) (int, error) {
			fmt.Println("Exec2", n)
			return 10, fmt.Errorf("error")
		},
		func(n int) int {
			fmt.Println("Rollback2", n)
			return n + 1
		},
	)

	result, err := tx.Process(1)
	fmt.Println(result, err)

}

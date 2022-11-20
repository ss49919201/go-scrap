package main

import (
	"fmt"
	"sync"
)

func main() {
	sm := sync.Map{}
	sm.Store("a", 1)
	fmt.Println(sm.Load("a"))
	sm.Delete("a")
	fmt.Println(sm.Load("a"))

	sm.Store(1, 1)
	fmt.Println(1, 1)

	sm.Store(struct{ int }{1}, 1)
	fmt.Println(sm.Load(struct{ int }{1}))

	i := 1
	i2 := 1
	sm.Store(struct{ *int }{&i}, 1)
	fmt.Println(sm.Load(struct{ *int }{&i2}))

	m := map[any]int{}
	m[1] = 1
	m["hoge"] = 1
	m[[]int{}] = 1
	fmt.Println(m)
}

package main

import (
	"encoding/json"
	"fmt"
)

type s struct {
	n  int
	np *int
}

func shallow(src *s) *s {
	return src
}

func deep(src *s) *s {
	b, _ := json.Marshal(src)
	var res *s
	json.Unmarshal(b, &res)
	return res
}

func main() {
	checkShallow()
	checkDeep()
}

func checkShallow() {
	a := new(s)
	b := new(s)
	a.n = 1
	b.n = 2
	fmt.Println(a == b)     // false
	fmt.Println(*a == *b)   // false
	fmt.Println(a.n == b.n) // false

	b = shallow(a)
	b.n = 3
	fmt.Println(a == b)     // true
	fmt.Println(*a == *b)   // true
	fmt.Println(a.n == b.n) // true}
}

func checkDeep() {
	a := new(s)
	b := new(s)
	a.n = 1
	b.n = 2
	fmt.Println(a == b)     // false
	fmt.Println(*a == *b)   // false
	fmt.Println(a.n == b.n) // false

	b = deep(a)
	b.n = 3
	fmt.Println(a == b)     // false
	fmt.Println(*a == *b)   // false
	fmt.Println(a.n == b.n) // false
}

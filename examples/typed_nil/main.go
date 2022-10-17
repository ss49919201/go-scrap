package main

import (
	"fmt"
)

type doI interface {
	do()
}

type product struct{}

func (p *product) do() {}

func a() *product {
	var p *product
	return p
}

func b() *product {
	return nil
}

// 型情報を持ったnil
func c() doI {
	var p *product
	return p
}

func d() doI {
	return nil
}

func main() {
	fmt.Println(a() == nil) // true
	fmt.Println(b() == nil) // true
	fmt.Println(c() == nil) // false
	fmt.Println(d() == nil) // true
}

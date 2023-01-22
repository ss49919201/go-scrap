package main

import (
	"fmt"

	"go.mercari.io/go-bps/bps"
)

func main() {
	n := bps.NewFromAmount(1)
	fmt.Println(n)
	fmt.Println(n.Amounts())
	fmt.Println(n.Percentages())
	fmt.Println(n.BasisPoints())
	fmt.Println(n.HalfBasisPoints())
	fmt.Println(n.DeciBasisPoints())
	fmt.Println(n.PPMs())
	fmt.Println(n.PPBs())
	// Output:
	// 100000
	// 1
	// 100
	// 10000
	// 20000
	// 100000
	// 1000000
	// 1000000000
}

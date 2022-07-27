package main

import (
	"fmt"

	"github.com/samber/mo"
)

func main() {
	var ip *int
	option1 := mo.Some(ip)
	fmt.Println(option1.OrEmpty())
}

package main

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
)

func main() {
	var ip *int
	option1 := mo.Some(ip)
	fmt.Println(option1.OrEmpty())                                                             // nil
	fmt.Println(option1.Get())                                                                 // nil, true
	fmt.Println(option1.IsPresent())                                                           // true
	fmt.Println(option1.IsAbsent())                                                            // false
	fmt.Println(option1.MustGet())                                                             // nil
	fmt.Println(option1.Size())                                                                // 1
	fmt.Println(option1.OrElse(lo.ToPtr(3)))                                                   // nil
	fmt.Println(option1.Map(func(value *int) (*int, bool) { return lo.ToPtr(2), true }).Get()) // 0xc0000b2020 true
	fmt.Println(option1.MapNone(func() (*int, bool) { return lo.ToPtr(2), true }).Get())       // nil true

	option1.ForEach(func(value *int) { fmt.Println(value) }) // nil

	fmt.Println(option1.FlatMap(func(value *int) mo.Option[*int] { return mo.None[*int]() }).Get()) // nil false

	onValue := func(value *int) (*int, bool) {
		return lo.ToPtr(2), true
	}
	onNone := func() (*int, bool) {
		return lo.ToPtr(4), true
	}
	option1.Match(onValue, onNone).ForEach(func(value *int) { fmt.Println(*value) }) // 2

	mo.None[*int]().Match(onValue, onNone).ForEach(func(value *int) { fmt.Println(*value) }) // 4
}

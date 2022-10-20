package main

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
)

func op() {
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

type user struct {
	id int
}

func getUser(id int) mo.Result[*user] {
	if id == 1 {
		return mo.Ok(&user{1})
	} else {
		return mo.Err[*user](errors.New("user not found"))
	}
}

func insertUser(u *user) mo.Result[*user] {
	if u.id == 1 {
		return mo.Ok(u)
	} else {
		return mo.Err[*user](errors.New("user not found"))
	}
}

func validateUser(u *user) mo.Result[*user] {
	return mo.Ok(u)
}

func createUser(u *user) mo.Result[*user] {
	return mo.Ok(u)
}

func main() {
	result1 := getUser(1).FlatMap(validateUser).FlatMap(createUser).FlatMap(insertUser)
	if result1.IsError() {
		fmt.Println(result1.Error())
	} else {
		fmt.Println(result1.MustGet())
	}

	result2 := getUser(2).FlatMap(validateUser).FlatMap(createUser).FlatMap(insertUser)
	if result2.IsError() {
		fmt.Println(result2.Error())
	} else {
		fmt.Println(result2.MustGet())
	}
}

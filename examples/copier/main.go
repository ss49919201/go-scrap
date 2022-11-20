package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type DefinedString string

type T struct {
	DS   DefinedString
	S    string
	PS   *string
	Time time.Time
}

func main() {
	dst := T{
		DS: DefinedString("ds"),
		S:  "s",
		PS: lo.ToPtr("ps"),
	}
	src := T{
		DS:   DefinedString("ds"),
		S:    "s",
		PS:   lo.ToPtr("ps"),
		Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	if err := copier.CopyWithOption(&dst, &src, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: copier.String,
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					_, ok := src.(string)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return "convert", nil
				},
			},
		},
	}); err != nil {
		panic(err)
	}

	fmt.Println("------1-------")
	rv := reflect.ValueOf(dst)
	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)
		if v.Kind() == reflect.Ptr {
			fmt.Println(v.Elem().String())
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println("------1-------")

	if err := copier.Copy(&dst, &src); err != nil {
		panic(err)
	}
	fmt.Println("------2-------")
	fmt.Println(reflect.DeepEqual(dst.DS, src.DS))
	fmt.Println(reflect.DeepEqual(dst.S, src.S))
	fmt.Println(reflect.DeepEqual(dst.PS, src.PS))
	fmt.Println(reflect.DeepEqual(dst.Time, src.Time))
	fmt.Println(dst.PS == src.PS)
	fmt.Println("------2-------")

	s := []int{1}
	var s2 []int
	if err := copier.Copy(&s2, &s); err != nil {
		panic(err)
	}
	fmt.Println("------3-------")
	fmt.Println(s2)
	fmt.Println(s[0] == s2[0])
	fmt.Println("------3-------")

	one := 1
	s3 := []*int{&one}
	var s4 []*int
	if err := copier.Copy(&s4, &s3); err != nil {
		panic(err)
	}
	fmt.Println("------4-------")
	fmt.Println(s3)
	fmt.Println(s3[0] == s4[0])
	fmt.Println(reflect.DeepEqual(s3[0], s4[0]))
	fmt.Println("------4-------")
	// Output:
	// ------1-------
	// ds
	// convert
	// convert
	// 2019-01-01 00:00:00 +0000 UTC
	// ------1-------
	// ------2-------
	// true
	// true
	// true
	// false
	// true
	// ------2-------
	// ------3-------
	// [1]
	// true
	// ------3-------
	// ------4-------
	// [0xc0000ac218]
	// false
	// true
	// ------4-------
}

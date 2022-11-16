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

	rv := reflect.ValueOf(dst)
	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)
		if v.Kind() == reflect.Ptr {
			fmt.Println(v.Elem().String())
		} else {
			fmt.Println(v)
		}
	}
	// Output:
	// ds
	// convert
	// convert
	// 2019-01-01 00:00:00 +0000 UTC
}

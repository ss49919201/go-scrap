- タグと値を取り出す
- TODO: フィールドがSliceや構造体の場合

```
package main

import (
	"fmt"
	"reflect"
)

type F struct {
	S  string  `json:"s"`
	I  int     `json:"i"`
	B  bool    `json:"b"`
	sp *string `json:"sp"`
}

func main() {
	var f F

	typ := reflect.TypeOf(f)
	fields := reflect.VisibleFields(typ)
	structElm := reflect.ValueOf(&f).Elem()

	for i, field := range fields {
		name := field.Name
		tag := field.Tag.Get("json")
		val := reflect.ValueOf(field)
		elm := structElm.Field(i)

		if elm.Kind() == reflect.Ptr {
			if elm.IsNil() {
				fmt.Println("🍙")
				fmt.Println("skip!")
				fmt.Printf("val: %s\n", val)
				fmt.Println("🍙")
				continue
			}
		}

		fmt.Println("--------------------")
		fmt.Printf("name: %s\n", name)
		fmt.Printf("tag: %s\n", tag)
		fmt.Printf("val: %s\n", val)
		fmt.Printf("elm: %v\n", elm)
		// fmt.Printf("elmInterface: %v\n", elmInterface)
		fmt.Println("--------------------")
	}
}
```
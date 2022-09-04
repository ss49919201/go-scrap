package pattern

import (
	"fmt"
	"reflect"
)

// pattern.ReflectValueOf("hoge")
// reflect.ValueOf(i): hoge
func ReflectValueOf(i interface{}) {
	fmt.Printf("reflect.ValueOf(i): %v", reflect.ValueOf(i))
}

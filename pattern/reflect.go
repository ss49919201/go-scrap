package pattern

import (
	"fmt"
	"reflect"
)

func ReflectValueOf(i interface{}) {
	fmt.Printf("reflect.ValueOf(i): %v", reflect.ValueOf(i))
}

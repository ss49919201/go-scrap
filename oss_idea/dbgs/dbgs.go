package dbgs

import (
	"reflect"
	"unsafe"
)

func SliceHeader[T any](s []T) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&s))
}

func SliceHeader2(s any) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&s))
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i, i2 *int
	fmt.Printf("%p,%p\n", i, i2)   // 0x0,0x0
	fmt.Printf("%p,%p\n", &i, &i2) // 0xc0000b2008,0xc0000b2010 (different addresses)
	var p1, p2 unsafe.Pointer = unsafe.Pointer(&i), unsafe.Pointer(&i2)
	fmt.Println(p1, p2) // 0xc0000b2008,0xc0000b2010 (different addresses)
}

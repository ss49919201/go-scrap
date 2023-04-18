package main

// #include <stdlib.h>
import (
	"C"
)

func main() {
	size := C.size_t(100)
	_ = C.malloc(size)
}

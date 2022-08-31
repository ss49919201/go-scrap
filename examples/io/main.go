package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	multiWrite()
	teeRead()
}

func teeRead() {
	r := bytes.NewBuffer([]byte("hello"))
	w := bytes.NewBuffer(nil)
	tr := io.TeeReader(r, w)
	tr.Read(make([]byte, 5))
	fmt.Println(w.String()) // hello
}

func multiWrite() {
	buf1 := bytes.NewBuffer(nil)
	buf2 := bytes.NewBuffer(nil)
	w := io.MultiWriter(buf1, buf2)
	w.Write([]byte("hello"))
	fmt.Println(buf1.String(), buf2.String()) // hello hello
}

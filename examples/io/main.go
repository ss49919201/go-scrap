package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {}

// 引数のbyte数だけバッファからコピーされる
func read() {
	buf := bytes.NewBuffer([]byte("hello!"))
	b := make([]byte, 5)
	buf.Read(b)
	fmt.Println(buf.String()) // !
	fmt.Println(string(b))    // hello
}

// 引数のbyteがバッファにコピーされる
func write() {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("hello"))
	fmt.Println(buf.String()) // hello
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

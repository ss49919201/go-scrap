package main

import (
	"fmt"
	"html/template"
)

func main() {
	if _, err := template.ParseFiles("examples/template/data/hello.tmpl"); err != nil {
		fmt.Println(err)
		return
	}

	// TODO:
	// fs := os.DirFS("data")
	// tmpl := template.New("hello")
	// _, err := tmpl.ParseFS(fs, "hello.tmpl")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

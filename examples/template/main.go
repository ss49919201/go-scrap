package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tmpl, err := template.ParseFiles("examples/template/data/hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := tmpl.Execute(os.Stdout, map[string]any{
		"Name": "beats",
		"Age":  20,
	}); err != nil {
		fmt.Println(err)
		return
	}
	// Hello beats
	// You are 20 years old.

	// TODO:
	// fs := os.DirFS("data")
	// tmpl := template.New("hello")
	// _, err := tmpl.ParseFS(fs, "hello.tmpl")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

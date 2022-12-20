package main

import (
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.md")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	result := markdown.ToHTML(b, nil, nil)
	f2, err := os.Create("test.html")
	if err != nil {
		log.Fatal(err)
	}
	f2.Write(result)
	f2.Close()
}

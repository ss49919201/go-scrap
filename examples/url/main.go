package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.JoinPath("http://example.com")
	fmt.Println(u) // http://example.com/

	u2, _ := url.JoinPath("http://example.com", "/")
	fmt.Println(u2) // http://example.com/

	u3, _ := url.JoinPath("http://example.com/", "/")
	fmt.Println(u3) // http://example.com/

	u4, _ := url.JoinPath("http://example.com//", "/")
	fmt.Println(u4) // http://example.com/

	u5, _ := joinPath("http://example.com//", "/")
	fmt.Println(u5) // http://example.com/
}

func joinPath(base string, elems ...string) (s string, err error) {
	u, err := url.Parse(base)
	if err != nil {
		return
	}
	s = u.JoinPath(elems...).String()
	return
}

package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func main() {
	u, _ := url.JoinPath("http://example.com/../go")
	p, _ := url.Parse("http://example.com/../go")
	fmt.Println(p.Path)
	fmt.Println(path.Join("../go/"))
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

// baseをurlとしてパースし、elemsをpathとしてurlに結合する
func joinPath(base string, elems ...string) (s string, err error) {
	u, err := url.Parse(base)
	if err != nil {
		return
	}
	s = u.JoinPath(elems...).String()
	return
}

func parse(u string) {
	// # より前と後ろに分ける
	_, _, _ = strings.Cut(u, "#")

	// # より前をパースする

	// # より後ろがあればそれをセットする、なければそのまま返す
}

// func joinPathToURL(u *url.URL, elem ...string) *url.URL {
// 	elem = append([]string{u.EscapedPath()}, elem...)
// 	var p string
// 	if !strings.HasPrefix(elem[0], "/") {
// 		// Return a relative path if u is relative,
// 		// but ensure that it contains no ../ elements.
// 		elem[0] = "/" + elem[0]
// 		p = path.Join(elem...)[1:]
// 	} else {
// 		p = path.Join(elem...)
// 	}
// 	// path.Join will remove any trailing slashes.
// 	// Preserve at least one.
// 	if strings.HasSuffix(elem[len(elem)-1], "/") && !strings.HasSuffix(p, "/") {
// 		p += "/"
// 	}
// 	url := *u
// 	url.setPath(p)
// 	return &url
// }

package pattern

import (
	"fmt"
	"strconv"
)

// s == {"id":1,"name":"hoge"} return "{\"id\":1,\"name\":\"hoge\"}"
func Escape(s string) string {
	return fmt.Sprintf("%q", s)
}

// s == `"{\"id\":1,\"name\":\"hoge\"}"` return {"id":1,"name":"hoge"}
// s == "{\"id\":1,\"name\":\"hoge\"}" invalid syntax
// s == "'a'" return "a"
// s == "'abc'" invalid syntax
func Unquote(s string) (string, error) {
	s, err := strconv.Unquote(s)
	if err != nil {
		return "", err
	}
	return s, nil
}

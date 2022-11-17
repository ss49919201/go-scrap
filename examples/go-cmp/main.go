package main

import (
	"time"

	"github.com/google/go-cmp/cmp"
)

func main() {
	type t struct {
		s string
		T time.Time
	}
	a := t{s: "a", T: time.Now()}
	b := t{s: "a", T: time.Now()}
	println(cmp.Equal(a, b), cmp.AllowUnexported(t{}, "s"))
	// println(cmp.Equal(a, b, cmpopts.IgnoreFields(t{}, "T")))
}

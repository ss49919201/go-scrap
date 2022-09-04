package pattern

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// go1.17

func Rune() {
	b := []rune{-214748358}
	println(string(b)) // �

	var buf bytes.Buffer
	buf.WriteRune(-214748358)
	println(buf.String())                   // �
	println(unicode.IsPrint(-214748358))    // false
	println(utf8.ValidString(buf.String())) // true
}

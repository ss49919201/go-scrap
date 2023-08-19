package main

import "github.com/thoas/go-funk"

type spf struct{}

func (_ spf) auth(input map[string]any) error {
	funk.Get(input, "envelopeFrom")
	return nil
}

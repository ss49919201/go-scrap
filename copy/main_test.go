package main

import "testing"

func TestShallow(t *testing.T) {
	a := "shallow"
	b := shallow(a)
	if a != b || &a != &b {
		t.Error("shallow is wrong")
	}
}

func TestDeep(t *testing.T) {
	a := "deep"
	b := deep(a)
	if a != b || &a == &b {
		t.Error("shallow is wrong")
	}
}

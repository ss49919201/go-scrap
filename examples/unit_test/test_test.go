package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("before")

	code := m.Run()

	println("after")

	os.Exit(code)
}

func Test_01(t *testing.T) {
	actual := "actual"
	expected := "expected"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_02(t *testing.T) {
	actual := "actual"
	expected := "expected"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

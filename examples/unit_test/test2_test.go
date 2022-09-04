package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_03(t *testing.T) {
	actual := "actual"
	expected := "expected"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_04(t *testing.T) {
	actual := "actual"
	expected := "expected"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_05(t *testing.T) {
	actual := "actual"
	expected := "expected"
	assert.Equal(t, actual, expected)
}

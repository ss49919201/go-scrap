package internal_test

import (
	"testing"

	"github.com/go-tips/examples/linq_generic/internal"
	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	result1 := internal.From([]string{"1", "2"}).ToSlice()
	assert.Equal(t, []string{"1", "2"}, result1)

	result2 := internal.From([]string{"1", "2"}).Where(func(s string) bool { return s == "2" }).ToSlice()
	assert.Equal(t, []string{"2"}, result2)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	assert.Equal(t, insert([]int{0, 1, 2}, 3, 3), []int{0, 1, 2, 3})
	assert.Equal(t, insert([]string{"0", "1", "2"}, 2, "1.5"), []string{"0", "1", "1.5", "2"})
}

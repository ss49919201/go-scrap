package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Run("merge sort", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, merge([]int{3, 1, 4, 5, 2}))
	})
}

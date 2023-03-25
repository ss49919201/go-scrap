package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("insert sort", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, insertionSort([]int{3, 1, 4, 5, 2}))
	})
}

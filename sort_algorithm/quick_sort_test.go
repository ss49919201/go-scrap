package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuick(t *testing.T) {
	t.Run("quick sort", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, quick([]int{3, 1, 4, 5, 2}))
	})
}

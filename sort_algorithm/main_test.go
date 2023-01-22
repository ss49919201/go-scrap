package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	t.Run("bubble sort", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, bubble([]int{3, 1, 4, 5, 2}))
	})
}

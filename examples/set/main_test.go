package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_base(t *testing.T) {
	set := newSet[string]()
	set.add("foo")
	set.add("bar")
	set.add("baz")
	set.add("foo")

	require.Equal(t, set.len(), 3)
	require.True(t, set.has("foo"))

	set.delete("foo")
	require.Equal(t, set.len(), 2)
	require.False(t, set.has("foo"))
}

package benchmark

import (
	"math"
	"testing"
)

var n = int(math.Pow10(6))

// 0.2221 ns/op
func BenchmarkMapStore(b *testing.B) {
	b.ResetTimer()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i
	}
}

// 0.01210 ns/op
func BenchmarkSliceAppend(b *testing.B) {
	b.ResetTimer()
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
}

// 0.001994 ns/op
func BenchmarkSliceBeforeInit(b *testing.B) {
	b.ResetTimer()
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = -1
	}

	for i := 0; i < n; i++ {
		s[i] = i
	}
}

package benchmark

import "testing"

// 0.2201 ns/op
func BenchmarkMapStore(b *testing.B) {
	b.ResetTimer()
	m := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
}

// 0.01251 ns/op
func BenchmarkSliceAppend(b *testing.B) {
	b.ResetTimer()
	s := []int{}
	for i := 0; i < 1000000; i++ {
		s = append(s, i)
	}
}

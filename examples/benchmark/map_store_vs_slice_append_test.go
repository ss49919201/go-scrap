package benchmark

import (
	"testing"
)

var n int = 1e6

func BenchmarkMapStore(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[i] = i
		}
	}
}

func BenchmarkSliceAppendA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{}
		for i := 0; i < n; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkSliceAppendB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, n)
		for i := 0; i < n; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkSliceBeforeInit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i
		}

		for i := 0; i < n; i++ {
			s[i] = i
		}
	}
}

package benchmark

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		s := "5"
		A(ss, s)
	}
}

func BenchmarkB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		s := "5"
		B(ss, s)
	}
}

// 要素がユニークであることが前提
func A(ss []string, s string) []string {
	r := make([]string, 0, len(ss)-1)
	for i := range ss {
		if ss[i] == s {
			r = append([]string{ss[i]}, r...)
			continue
		}
		r = append(r, ss[i])
	}
	return r
}

// 要素がユニークであることが前提
func B(ss []string, s string) []string {
	for i := 0; i < len(ss); i++ {
		if ss[i] == s {
			v := ss[i]
			ss = append([]string{v}, append(ss[:i], ss[i+1:]...)...)
		}
	}
	return ss
}

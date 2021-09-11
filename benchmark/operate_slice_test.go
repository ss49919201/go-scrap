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

func BenchmarkC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		s := "5"
		C(ss, s)
	}
}

func BenchmarkD(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		s := "5"
		D(ss, s)
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
	for i := range ss {
		if ss[i] == s {
			v := ss[i]
			ss = append([]string{v}, append(ss[:i], ss[i+1:]...)...)
		}
	}
	return ss
}

func C(ss []string, s string) []string {
	for i := range ss {
		if ss[i] == s {
			v := ss[i]
			ss = append(ss[:i], ss[i+1:]...)
			ss = append([]string{v}, ss...)
		}
	}
	return ss
}

func D(ss []string, s string) []string {
	for i, v := range ss {
		if ss[i] == s {
			ss = append([]string{v}, append(ss[:i], ss[i+1:]...)...)
		}
	}
	return ss
}

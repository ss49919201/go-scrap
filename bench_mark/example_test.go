package main

import (
	"math/rand"
	"testing"
)

var n int = 1000
var testcaseA []int
var testcaseB []int

func init() {
	testcaseA = func() []int {
		r := []int{}
		for i := 1; i <= n; i++ {
			r = append(r, rand.Intn(n))
		}
		return r
	}()

	testcaseB = func() []int {
		r := []int{}
		for i := 1; i <= n; i++ {
			r = append(r, rand.Intn(n))
		}
		return r
	}()
}

func Benchmark_linearSearch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linearSearch(testcaseA, testcaseB)
	}
}

func Benchmark_inearSearchWithHashMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linearSearchWithHashMap(testcaseA, testcaseB)
	}
}

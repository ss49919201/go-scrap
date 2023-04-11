package benchmark

import (
	"math/rand"
	"testing"
)

var n int = 1e1

var testcaseA = func() []int {
	r := []int{}
	for i := 1; i <= n; i++ {
		r = append(r, rand.Intn(n))
	}
	return r
}()

var testcaseB = func() []int {
	r := []int{}
	for i := 1; i <= n; i++ {
		r = append(r, rand.Intn(n))
	}
	return r
}()

func Benchmark_linearSearch(t *testing.B) {
	linearSearch(testcaseA, testcaseB)
}

func Benchmark_inearSearchWithHashMap(t *testing.B) {
	linearSearchWithHashMap(testcaseA, testcaseB)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var number int = 1000
var caseA []int
var caseB []int

func init() {
	caseA = func() []int {
		r := []int{}
		for i := 1; i <= number; i++ {
			r = append(r, rand.Intn(number))
		}
		return r
	}()

	caseB = func() []int {
		r := []int{}
		for i := 1; i <= number; i++ {
			r = append(r, rand.Intn(number))
		}
		return r
	}()
}

func main() {
	// fmt.Println(caseA, caseB)
	start := time.Now()
	// linearSearchWithHashMap(caseA, caseB)
	linearSearch(caseA, caseB)
	fmt.Println(time.Now().Sub(start))
}

func linearSearch(a, b []int) {
	var result []int
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				result = append(result, v)
				break
			}
		}
	}
}

func linearSearchWithHashMap(a, b []int) {
	var result []int

	hashmap := make(map[int]struct{}, len(b))
	for _, v := range b {
		hashmap[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := hashmap[v]; ok {
			result = append(result, v)
		}
	}
}

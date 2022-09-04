package main

import (
	"math/rand"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	ss1 := generateRandomUniqStrings()
	ss2 := generateRandomUniqStrings()

	// 同じ要素が含まれていることを担保
	if !lo.Every(ss1, ss2) {
		panic("ss1 and ss2 must have same elements")
	}

	ss3 := mergeAndSort(ss1, ss2, 10)
	ss4 := sortAndMerge(ss1, ss2, 10)

	// 両方の結果が同じであることを担保
	lo.ForEach(ss3, func(s string, i int) {
		if s != ss4[i] {
			panic("ss3 and ss4 must have same elements and order")
		}
	})
}

// 毎回同じ乱数を生成するためにシードを固定
// 実行する毎に異なるが必要順序な場合は`rand.Seed(time.Now().UnixNano())`
func generateRandomUniqStrings() []string {
	ss := make([]string, 0, 100)
	for i := 1; i <= 100; {
		s := strconv.Itoa(rand.Intn(100))
		if lo.Contains(ss, s) {
			continue
		}
		ss = append(ss, s)
		i++
	}
	return ss
}

func mergeAndSort(ss1, ss2 []string, limit int) []string {
	ss := make([]string, 0, len(ss1)+len(ss2))
	ss = append(ss, ss1...)
	ss = append(ss, ss2...)
	sort.Strings(ss)
	return ss[:limit]
}

func sortAndMerge(ss1, ss2 []string, limit int) []string {
	sort.Strings(ss1)
	sort.Strings(ss2)
	ss := make([]string, 0, limit*2)
	ss = append(ss, ss1[:limit]...)
	ss = append(ss, ss2[:limit]...)
	sort.Strings(ss)
	return ss[:limit]
}

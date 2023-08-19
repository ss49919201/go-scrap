package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"apple":  5,
		"banana": 3,
		"fruits": map[string]interface{}{
			"orange": 8,
			"pear":   2,
		},
	}

	printTable(data, 0)
}

func printTable(data map[string]interface{}, level int) {
	// キーを取得し、ソートする
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 最長のキー長を計算する
	maxKeyLen := 0
	for _, key := range keys {
		keyLen := len(key)
		if keyLen > maxKeyLen {
			maxKeyLen = keyLen
		}
	}

	// テーブルを出力する
	fmt.Println(strings.Repeat("-", maxKeyLen+9+(level*2)))
	fmt.Printf("| %-*s | %*svalue |\n", maxKeyLen, "key", level*2, "")
	fmt.Println(strings.Repeat("-", maxKeyLen+9+(level*2)))
	for _, key := range keys {
		switch v := data[key].(type) {
		case map[string]interface{}:
			fmt.Printf("| %-*s | %*s=> map |\n", maxKeyLen, key, level*2, "")
			printTable(v, level+1)
		default:
			fmt.Printf("| %-*s | %*s%5v |\n", maxKeyLen, key, level*2, "", data[key])
		}
	}
	fmt.Println(strings.Repeat("-", maxKeyLen+9+(level*2)))
}

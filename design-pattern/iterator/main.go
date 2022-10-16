package main

import "unicode/utf8"

// 複数要素を持つ構造体の各要素へアクセスする方法を定義
// 配列がmapになったり内部構造が変わっても呼び出し側に影響がない
// 逆から呼び出すとか特定の条件の要素はアクセスさせないとかの変更がしやすい？この場合は実装を変える？
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func NewIteratorA() Iterator[string] {
	return &IteratorImplA{
		elements: []string{"a", "b", "c"},
	}
}

type IteratorImplA struct {
	elements     []string
	currentIndex int
}

func (i *IteratorImplA) HasNext() bool {
	return i.currentIndex < len(i.elements)
}

func (i *IteratorImplA) Next() string {
	defer func() {
		i.currentIndex++
	}()

	return i.elements[i.currentIndex]
}

func NewIteratorB() Iterator[int] {
	return &IteratorImplB{
		elements:   map[string]int{"a": 0, "b": 1, "c": 2},
		currentKey: "a",
	}
}

type IteratorImplB struct {
	elements   map[string]int
	currentKey string
}

func (i *IteratorImplB) HasNext() bool {
	c, _ := utf8.DecodeRuneInString("c")
	r, _ := utf8.DecodeRuneInString(i.currentKey)
	return r <= c
}

func (i *IteratorImplB) Next() int {
	defer func() {
		k, _ := utf8.DecodeRuneInString(i.currentKey)
		i.currentKey = string(k + 1)
	}()

	return i.elements[i.currentKey]
}

func main() {
	ia := NewIteratorA()
	for ia.HasNext() {
		println(ia.Next())
	}

	ib := NewIteratorB()
	for ib.HasNext() {
		println(ib.Next())
	}
}

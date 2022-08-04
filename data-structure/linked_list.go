package main

// 要素とポインタからなる数珠繋ぎのリスト
// 連続した領域にならないので事前にメモリ確保が不要
// 頻繁にデータを書き換えるのならば有効
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *List[T]) Add() {
	// TODO
}

func (l *List[T]) Remove() {
	// TODO
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func main() {
	_ = new(List[int])
}

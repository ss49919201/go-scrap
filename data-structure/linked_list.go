package main

// 要素とポインタからなる数珠繋ぎのリスト
// 連続した領域にならないので事前にメモリ確保が不要
// 頻繁にデータを書き換えるのならば有効

type ListIF interface {
	// 先頭に要素を追加する
	PushHead(e Node[any])
	// 末尾に要素を追加する
	PushTail(e Node[any])
	// 先頭の要素を取り出す
	PopHead() Node[any]
	// 末尾の要素を取り出す
	PopTail() Node[any]
	// 要素数を取得する
	Len() int
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func main() {
	_ = new(List[int])
}

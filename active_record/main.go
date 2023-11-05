package main

import "time"

var books = map[int]Book{}

// データと振る舞いを持つオブジェクト
// レコードをラップしている
type Book struct {
	// データの「ほとんど」が永続化される
	ID         int
	Title      string
	Label      int
	PublishdAt time.Time
}

// 永続化されたデータを操作するロジックをオブジェクトの振る舞いとして実装
// ※本来DBが永続化領域になるが、mapで代用している
func (b Book) Save() {
	books[b.ID] = b
}

func (b *Book) Delete() {
	delete(books, b.ID)
}

// ドメインロジックをオブジェクトの振る舞いとして実装
func (b *Book) IsLatest() bool {
	return b.PublishdAt.After(time.Now().AddDate(0, -1, 0))
}

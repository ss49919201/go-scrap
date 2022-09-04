package main

import "fmt"

// http://objectclub.jp/community/memorial/homepage3.nifty.com/masarl/article/dp-ocp-2.html
// 仕様変更時にコードの修正ではなく追加で対応することで、影響範囲を狭くする
// 多態性を持たせる

// Open
// 振る舞いを拡張する際には、修正は不要で追加のみ
// 変更が多そうなモジュールはinterfaceにして、仕様追加時には実装を増やす
// 既存の呼び出し元や既存の実装を変更しないで、新たな実装を増やせる

// Closed
// 仕様変更時には依存しているモジュールは変更しなくて良い
// 修正箇所が対象モジュールのみに閉じている

// - Repositoryをinterfaceにする
// 	- 呼び出し元がユニットテストの場合は、interfaceを実装しているMockを使用できる
// - Nodeをinterfaceにする
// 	- 全てのNodeをLoopで回して...といった処理で、新たなNode実装を増やす以外は変更が不要になる

// 解決策として一つのデザインパターンがあるわけではない
// 変更されやすい箇所に応じてデザインパターンを使い分ける

// ----------------------------------------------------------------------

type UserImpl1 struct{}

func (u *UserImpl1) Find(id int) {
	fmt.Println("UserImpl1.Find")
}

// 基本のリポジトリ
type UserRepo interface {
	Find(id int)
}

// 拡張リポジトリを実装する
// open : 拡張したいメソッドを持つinterfaceを追加で定義することで、既存の機能を損なうことなく拡張可能
// closed : 追加のみで対応できるので、基本のリポジトリの修正がいらない
type UserRepoExt interface {
	UserRepo
	FindByName(name string)
}

type UserImpl2 struct {
	UserImpl1
}

func (u *UserImpl2) Find(id int) {
	fmt.Println("UserImpl2.Find")
}

func (u *UserImpl2) FindByName(name string) {}

func phase1() {
	var repo UserRepo = &UserImpl1{}
	repo.Find(1)
}

func phase2() {
	var repo UserRepoExt = &UserImpl2{}
	repo.Find(1)
	repo.FindByName("")
}

func main() {
	phase1()
	phase2()
}

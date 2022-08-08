package main

// http://objectclub.jp/community/memorial/homepage3.nifty.com/masarl/article/dp-ocp-2.html
// 仕様変更時にコードの修正ではなく追加で対応することで、影響範囲を狭くする
// 多態性を持たせる
// 振る舞いを拡張する際には、修正は不要で追加のみ
// 変更が多そうなモジュールはinterfaceにして、仕様追加時には実装を増やす
// 既存の呼び出し元や既存の実装を変更しないで、新たな実装を増やせる
// - Repositoryをinterfaceにする
// 	- 呼び出し元がユニットテストの場合は、interfaceを実装しているMockを使用できる
// - Nodeをinterfaceにする
// 	- 全てのNodeをLoopで回して...といった処理で、新たなNode実装を増やす以外は変更が不要になる

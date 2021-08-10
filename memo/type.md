# type

# type alias とは

- 既存の型に紐付く型を宣言すること
    - 宣言された型は、紐付く型であって新たに定義された型ではない
		- defeined type とは異なる
- 既存の型名を宣言しようとすると当然同ブロック内での際宣言なのでコンパイルエラー`redeclared in this block`となる
- refactorに用いる為にできた仕様

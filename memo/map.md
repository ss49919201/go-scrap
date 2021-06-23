# map

## mapとは

- ハッシュテーブルへの参照
    - ハッシュテーブルとは順番の決まっていないKeyとValueの集合を表すデータ構造
    > ハッシュテーブルは連想配列や集合の効率的な実装のうち1つである。(https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB)
- キーの型は`==`で比較可能である必要がある
    - 関数は`==`で比較できない
    - [Goで比較可能な型](https://golang.org/ref/spec#Comparison_operators)
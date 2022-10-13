package main

import "fmt"

var flag = make(map[int]bool)
var e = make(map[int][]int)
var deque []int
var stop bool

// https://atcoder.jp/contests/abc270/editorial/4878
// > 1.duque の末尾に v を追加する。
// 2.現在いる頂点 v について、flag[v]=True に変更する。
// 3.v から移動できる頂点 w であってまだ訪ねていないもの、すなわち flag[w]=False であるようなものを選び、頂点 w に移動した後、 1.から再度行う。
// 4.duque の末尾から要素を 1 つ削除する。（ここで、削除される要素は v 自身となる）
// 5.3.においてそのような頂点が存在しない場合、現在いる頂点がスタートした頂点 X であるならば探索を終了する。 そうでない時、この頂点を始めて訪れる直前のマスに戻り、2.の作業を再開する。

func dfs(cur, next int) {
	// duque の末尾に cur を追加
	if stop {
		return
	} else {
		deque = append(deque, cur)
	}

	// 目的地に到達したら終了
	if cur == next {
		stop = true
		return
	}

	// 頂点到達フラグを更新
	flag[cur] = true

	// 現在地から行ける頂点を探索
	// 既に到達済みの頂点は探索しない
	for _, v := range e[cur] {
		if !flag[v] {
			dfs(v, next)
		}
	}

	// 現在地から行ける頂点全てが訪問済み又は行き止まりの場合は
	// 末尾の頂点を削除 (つまり、一つ前の頂点での探索に戻る)
	if stop {
		return
	} else {
		popBack(&deque)
	}
}

func popBack(l *[]int) int {
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e
}

// https://atcoder.jp/contests/abc270/tasks/abc270_c
func main() {
	X, Y := 2, 5

	ls := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	for _, v := range ls {
		e[v[0]] = append(e[v[0]], v[1])
		e[v[1]] = append(e[v[1]], v[0])
	}

	dfs(X, Y)

	for i, v := range deque {
		fmt.Print(v)
		if (i + 1) != len(deque) {
			fmt.Print(" ")
		}
	}

	// Output:
	// 2 1 3 5
}

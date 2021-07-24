# defer

# deferとは

- Goのおけるstatement(文)の一つ
- deferへは関数を渡し、渡した関数は呼び出し元の関数がreturnする直前に実行

> A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.

https://golang.org/ref/spec#Defer_statements

- LIFO
- 組み込み関数は使えないものが多い

> The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for expression statements.

https://golang.org/ref/spec#Defer_statements

# サンプル

```
func callDefer() {
	// fmt.Println(bigSlowOperation())
	defer f()()
	print(2)
}

func f() func() {
	print(1)
	return func() { print(3) }
}
```

1. f()の実行により1出力
2. print(2)により2出力
3.f()の返り値の関数実行により3出力

```
func callDefer() (i int){
	// fmt.Println(bigSlowOperation())
	defer func(){
		i = 4
	}()
	defer f()()
	print(2)
	return 0
}

func f() func() {
	print(1)
	return func() { print(3) }
}
```

1. ()の実行により1出力
2. print(2)により2出力
3. sに0の代入
4. f()の返り値の関数実行により3出力
5. sに4の代入

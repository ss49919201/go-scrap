# Slice

## スライスのポインタ型

スライスのポインタ型はインデックス操作ができない。

```go
s := &[]int{1, 2}
fmt.Println(*s[0])
```

```sh
./tmp.go:13:16: invalid operation: s[0] (type *[]int does not support indexing)
```
## スライス演算子

スライス演算子を用いてスライスを生成する際に指定できるインデックスの範囲と、インデックス指定により要素にアクセスする際に指定できる範囲は異なる。

> the index x is in range if 0 <= x < len(a), otherwise it is out of range

https://golang.org/ref/spec#Index_expressions

> For arrays or strings, the indices are in range if 0 <= low <= high <= len(a), otherwise they are out of range. 

https://golang.org/ref/spec#Slice_expressions

```go
	s := []int{1, 2}
    // 要素へのアクセス
	fmt.Println(s[1])  // 2
	fmt.Println(s[2])  // panic: runtime error: index out of range [2] with length 2
    // スライスの生成
	fmt.Println(s[1:]) // [2]
	fmt.Println(s[:2]) // [1 2]
	fmt.Println(s[2:]) // [2]
```
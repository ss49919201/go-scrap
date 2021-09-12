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

https://golang.org/ref/spec#Slice_expressions
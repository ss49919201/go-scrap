# ジェネリクスについて

- `go run -gcflags=-G=3 generics.go`

## メリット

- ゼロ値を返せる

```go
func gZero[T any](s T) {
	var zero T
	fmt.Println(zero)
}

func iZero(s interface{}) {
	var zero interface{}
	fmt.Println(zero)
}

func main() {
	var i = 1
	gZero(i)
	iZero(i)
	var s = "hoge"
	gZero(s)
	iZero(s)
}
```

```sh
0
<nil>

<nil>
```
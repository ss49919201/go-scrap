- interfaceは埋め込める

```
type s interface {
	fmt.Stringer
}
```

- interface型に型アサーション可能

```
type hoge struct{}

func (h hoge) String() string { return "hoge" }

func (h hoge) fuga() string { return "fuga" }

type fuga interface {
	fuga() string
}

func main() {
	var s hoge
	var i fmt.Stringer
	i = s
	if _, ok := i.(fuga); true {
		fmt.Println(ok) // true
	}

	if _, ok := i.(hoge); true {
		fmt.Println(ok) // true
	}
}

```
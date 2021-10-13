```
slice := []string{"hoge", "fuga", "piyo"}

func(ps []string) {
	fmt.Println(fmt.Sprintf("%p\n", ps) == fmt.Sprintf("%p\n", slice)) // true
}(slice)
func(ps ...string) {
	fmt.Println(fmt.Sprintf("%p\n", ps) == fmt.Sprintf("%p\n", slice)) // true
}(slice...)
```
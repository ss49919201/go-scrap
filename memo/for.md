第二引数を省略するとtrueで評価される
https://golang.org/ref/spec#For_statements

```
func main() {
	for i := 0; ; i++ {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}
```
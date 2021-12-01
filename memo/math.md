## 小数点以下の切り上げ

```go
func main() {
	fmt.Println(math.Ceil(1.0))     // 1
	fmt.Println(math.Ceil(1.01))    // 2
	fmt.Println(math.Ceil(0.9))     // 1
	fmt.Println(math.Ceil(0.00001)) // 1
}
```

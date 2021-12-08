## 小数点以下の切り上げ

```go
func main() {
	fmt.Println(math.Ceil(1.0))     // 1
	fmt.Println(math.Ceil(1.01))    // 2
	fmt.Println(math.Ceil(0.9))     // 1
	fmt.Println(math.Ceil(0.00001)) // 1
	fmt.Println(math.Ceil(-3.0))    // -3

	fmt.Println(math.Ceil(-0.1))      // -0 符号がつくことに注意
	fmt.Println(math.Ceil(-0.1) == 0) // true -0と0の比較はtrue

	strNegZero := strconv.Itoa(int(math.Ceil(-0.1)))
	strZero := strconv.Itoa(0)
	fmt.Printf("compare %s, %s\n", strNegZero, strZero) // compare 0, 0
	fmt.Println(strNegZero == strZero)                  // true -0と0の文字列変換は等しく0
}
```

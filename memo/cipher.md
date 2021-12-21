https://deeeet.com/writing/2015/11/10/go-crypto/

```go
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	// ブロック暗号化
	text := []byte("1234567890abcdef")
	// 16byte
	key := []byte("1234567890abcdef")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cipher := make([]byte, len(text))
	block.Encrypt(cipher, text)
	fmt.Println(string(text))   // 1234567890abcdef
	fmt.Println(string(cipher)) // ������\7��e�c_
}
```

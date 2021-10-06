```
package main

import (
	"fmt"
)

func p() {
	panic("panic!")
}

func main() {
	fmt.Println("start")
	p() // endは出力されない
	fmt.Println("end")
}

```
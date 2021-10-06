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

```
package main

import (
	"fmt"
)

func np() {
	defer fmt.Println("defer np") // 出力される
	panic("panic!")
}

func p() {
	defer fmt.Println("defer p") // 出力される
	np()
	panic("panic!")
}

func main() {
	fmt.Println("start")
	defer fmt.Println("defer main") // 出力される
	p()
	fmt.Println("end") // 出力されない
}

```

```
package main

import (
	"fmt"
)

// panic! np, defer panic np, deffer panic p, defer panic main の順でpanicが呼ばれる

func np() {
	defer panic("defer panic np") // 出力される
	panic("panic! np")
}

func p() {
	defer panic("defer panic p") // 出力される
	np()
	panic("panic! p")
}

func main() {
	fmt.Println("start")
	defer panic("defer panic main") // 出力される
	p()
	fmt.Println("end") // 出力されない
}

```
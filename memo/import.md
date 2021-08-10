# import

# import宣言について

- package宣言の後にimport宣言を書ける
    - package宣言の前には書けない
        - package宣言は必ずソースファイルの先頭
        - 複数のpackage宣言もできない
```go
import "fmt" // expected 'package', found'import'
package main
```

```go
package main
package hoge // syntaxerror:non-declaration statementoutsidefunction body
// package宣言はnon-declaration?
```

# blank import

- ブランク識別子`_`を用いた改名import
    - ブランク識別子は参照できない
    - `unused import`を防げる
    - そのpackageのimportによる副作用を利用する為に 用いる
        - packageレベルの変数初期化、init関数等

```go
import (
    _ "fmt"
)
```

# dot import

- importしたpackageの、packageブロックで宣言された識別子を修飾子なしで参照できる
- `should not use dot imports (ST1001)go-staticcheck`

```go
package main

import (
	. "fmt"
)

func main() {
	Println()
}
```
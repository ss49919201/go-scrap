# string

## string とは

- 文字列を示す型
- 文字列はコンピュータから見れば、0と1(ビット)の列
- Goの文字列はイミュータブルなバイト列
    - 文字列を構成するバイトは0から255までの符号なしの8ビットの整数

```go
s := "文字列🍙"
fmt.Println(len(s)) // 13 つまり文字列は13バイトから成る配列
``` 

- 文字列を構成しているバイトにもアクセス可能

```go
s := "文字列🍙"
fmt.Println(s[0]) // 230
```

- `[]byte`にも`[]rune`にも変換可能
    - runeは文字のUnicodeコードポイントを示す整数値
    - runeはint32の型alias
    - byteはuint8の型alias
        - 型aliasはスコープ内であれば同一の型として扱えるということ
- 「文」のUnicodeコードポイントは6587(10進数だと25991)
    - UTF-8でエンコードするとE69687(2進数だと111001101001011010000111)
    - 111001101001011010000111を先頭から1バイトごとに区切ると、[11100110,10010110,10000111](10進数だと[230,150,135])
- 「字」のUnicodeコードポイントは5B57(10進数だと23383)
- 「列」のUnicodeコードポイントは5217(10進数だと21015)
- 「🍙」のUnicodeコードポイントは1F359(10進数だと127833)

* 「文」以外のコードポイントエンコード結果の記載は省略

```go
fmt.Println([]byte(s)) // [230 150 135 229 173 151 229 136 151 240 159 141 153]
fmt.Println([]rune(s)) // [25991 23383 21015 127833]

[]byte(s)[0] == s[0] // true
```

- 文字列も範囲式として指定できる
- ループの対象となる要素はUnicodeコードポイント
    - 第二値はコードポイント(rune)となる
    - 第一値はコードポイントのインデックスではなく、UTF-8エンコードされたバイト列の先頭のインデックスになるので注意

```go
for i, v := range s {
		fmt.Println("s["+strconv.Itoa(i)+"]", s[i])
		fmt.Println("v", v)
		/*
		s[0] 230
		v 25991
		s[3] 229
		v 23383
		s[6] 229
		v 21015
		s[9] 240
		v 127833
		*/
	}
```

# 参考

- http://ash.jp/code/unitbl21.htm
- https://golang.org/ref/spec#For_statements
- https://0g0.org/unicode/1F359/
- https://hogehoge.tk/tool/number.html# string
# パース処理

# 整数から文字列へのパース

```go
func ParseInt() {
	// baseには基数(0,2~36)を指定
	// bitsizeにはビット長(8~64)を指定
	// 引数に応じた数値がint64で返ってくる
	var i uint8 = math.MaxUint8
	print(i, " ", i+1, "\n")
	print((^uint(0) >> 63), "\n")
	print(math.MaxInt32, strconv.Itoa(math.MinInt32+1), "\n")
	v32 := "-354634382"
	// bitsizeが0の場合は関数内で64に変換される
	if s, err := strconv.ParseInt(v32, 10, 0); err == nil {
		fmt.Printf("strconv.ParseInt(v32, 10, 32) %T, %v\n", s, s)
	} else {
		fmt.Printf("strconv.ParseInt(v32, 10, 32) %T, %v\n", err, err)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("strconv.ParseInt(v32, 16, 32) %T, %v\n", s, s)
	} else {
		fmt.Printf("strconv.ParseInt(v32, 16, 32) %T, %v\n", err, err)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("strconv.ParseInt(v64, 10, 64) %T, %v\n", s, s)
	} else {
		fmt.Printf("strconv.ParseInt(v64, 10, 64) %T, %v\n", err, err)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("strconv.ParseInt(v64, 16, 64) %T, %v\n", s, s)
	} else {
		fmt.Printf("strconv.ParseInt(v64, 16, 64) %T, %v\n", err, err)
	}
}

```
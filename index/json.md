```go
var J j
b, _ := json.Marshal(J)
println(string(b))
```

```go
var J2 *j
err := json.Unmarshal(b, J2) // pointerじゃないとダメ。nil（pointer型のゼロ値）はダメ。
if err != nil {
	fmt.Println(err.Error())
	os.Exit(1)
}
println(string(b))
```
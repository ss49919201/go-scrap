```
a := []string{0: "hello", 1: "world"}
b := []string{0: "hello", 1: "world"}
// fmt.Println(a == b) // 比較不可
fmt.Println(reflect.DeepEqual(a, b)) // true
```

```
s := ""
s2 := ""
var a struct {
	a *string
	b string
} = struct {
	a *string
	b string
}{a: &s}
var b struct {
	a *string
	b string
} = struct {
	a *string
	b string
}{a: &s2}
fmt.Println(a == b)                  // false
fmt.Println(reflect.DeepEqual(a, b)) // true
```

```
s := ""
s2 := ""
var a struct {
	a *string
	b string
} = struct {
	a *string
	b string
}{a: &s}
var b *struct {
	a *string
	b string
} = &struct {
	a *string
	b string
}{a: &s2}
// fmt.Println(a == b)                  // 比較不可
fmt.Println(reflect.DeepEqual(a, b)) // false
```
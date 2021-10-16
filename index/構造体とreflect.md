- タグと値を取り出す
- TODO: フィールドがSliceや構造体の場合

```
type F struct {
	S  string  `json:"s"`
	I  int     `json:"i"`
	B  bool    `json:"b"`
	sp *string `json:"sp"`
}

func main() {
	var f F

	typ := reflect.TypeOf(f)
	fields := reflect.VisibleFields(typ)
	structElm := reflect.ValueOf(&f).Elem()

	for i, field := range fields {
		name := field.Name
		tag := field.Tag.Get("json")
		val := reflect.ValueOf(field)
		elm := structElm.Field(i)

		if elm.Kind() == reflect.Ptr {
			if elm.IsNil() {
				fmt.Println("\n🍙")
				fmt.Println("skip!")
				fmt.Printf("val: %s\n", val)
				fmt.Println("🍙")
				continue
			}
		}

		fmt.Println("\n--------------------")
		fmt.Printf("name: %s\n", name)
		fmt.Printf("tag: %s\n", tag)
		fmt.Printf("val: %s\n", val)
		fmt.Printf("elm: %v\n", elm)
		fmt.Println("--------------------")
	}
}

output:

--------------------
name: S
tag: s
val: {S  string json:"s" %!s(uintptr=0) [%!s(int=0)] %!s(bool=false)}
elm: 
--------------------

--------------------
name: I
tag: i
val: {I  int json:"i" %!s(uintptr=16) [%!s(int=1)] %!s(bool=false)}
elm: 0
--------------------

--------------------
name: B
tag: b
val: {B  bool json:"b" %!s(uintptr=24) [%!s(int=2)] %!s(bool=false)}
elm: false
--------------------
🍙
skip!
val: {sp main *string json:"sp" %!s(uintptr=32) [%!s(int=3)] %!s(bool=false)}
🍙
~/src/go-tips ● ▲ ■  go run main.go                                                                                          (git)-[master]

--------------------
name: S
tag: s
val: {S  string json:"s" %!s(uintptr=0) [%!s(int=0)] %!s(bool=false)}
elm: 
--------------------

--------------------
name: I
tag: i
val: {I  int json:"i" %!s(uintptr=16) [%!s(int=1)] %!s(bool=false)}
elm: 0
--------------------

--------------------
name: B
tag: b
val: {B  bool json:"b" %!s(uintptr=24) [%!s(int=2)] %!s(bool=false)}
elm: false
--------------------

🍙
skip!
val: {sp main *string json:"sp" %!s(uintptr=32) [%!s(int=3)] %!s(bool=false)}
🍙
```
```
	var a []*string
	fmt.Printf("%#v\n", a) // []*string(nil)
	fmt.Println(len(a))    // 0
	fmt.Println(nil == a)  // true

	b := []*string{}
	fmt.Printf("%#v\n", b) // []*string{}
	fmt.Println(len(b))    // 0
	fmt.Println(nil == b)  // false

	c := []*string{nil}
	fmt.Printf("%#v\n", c) // []*string{(*string)(nil)}
	fmt.Println(len(c))    // 1
	fmt.Println(nil == c)  // false

	fmt.Println(reflect.DeepEqual(a, b)) // false
	fmt.Println(reflect.DeepEqual(b, c)) // false
```

```
	a := []string{0: "hello", 1: "world"}
	fmt.Println(len(a)) // 2
	fmt.Println(a)      // [hello world]

	b := []string{0: "hello", 2: "world"}
	fmt.Println(len(b)) // 3
	fmt.Println(b)      // [hello  world]
```
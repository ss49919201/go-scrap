FIXME: 解説
```
type e struct{}

func (e *e) Error() string { return "e" }

func ChAs() {
	// err := fmt.Errorf("%w", errors.New(""))
	err := fmt.Errorf("%w", &e{})
	var target interface{}
	var se *e
	target = &se
	fmt.Println(errors.As(err, target))
	var ie = &e{}
	var i interface{} = &ie
	println(reflect.DeepEqual(i, target))
}

```
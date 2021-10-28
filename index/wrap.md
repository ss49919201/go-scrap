```go
we := fmt.Errorf("%w", errors.New("error"))
v := reflect.ValueOf(we)
fmt.Println(v.Type().String()) // *fmt.wrapError
pp.Println(we)
// &fmt.wrapError{
//   msg: "error",
//   err: &errors.errorString{
//     s: "error",
//   },
// }

// 既にラップ済みのものを引数にしたら？
wwe := fmt.Errorf("%w", we)
v = reflect.ValueOf(wwe)
fmt.Println(v.Type().String()) // *fmt.wrapError
pp.Println(wwe)
// 	&fmt.wrapError{
//   msg: "error",
//   err: &fmt.wrapError{
//     msg: "error",
//     err: &errors.errorString{
//       s: "error",
//     },
//   },
// }
```

```go
we := fmt.Errorf("%w", errors.New("error"))
uwe := errors.Unwrap(we)
uwe2 := errors.Unwrap(errors.New("unwrap"))
pp.Println(uwe)
pp.Println(uwe2)
// 	&errors.errorString{
//   s: "error",
// }
// nil
```

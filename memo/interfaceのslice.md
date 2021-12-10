```
func main() {
	slice := []struct {
		i int
		s string
	}{
		{
			i: 1,
			s: "1",
		}, {
			i: 2,
			s: "2",
		}, {
			i: 3,
			s: "3",
		},
	}
	iSlice := make([]interface{}, len(slice))
	for i := range slice {
		iSlice[i] = interface{}(slice[i])
	}
	run(iSlice...)
}

func run(iSlice ...interface{}) {
	for i := range iSlice {
		fmt.Println(iSlice[i])
	}
}
```
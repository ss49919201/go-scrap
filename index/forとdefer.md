- createが10個出力されてから、deferが出力される

```
func main() {
	run()
}

func run() {
	for i := 0; i < 10; i++ {
		s := create()
		defer delete(s)
	}
}

func create() string {
	now := time.Now()
	u, _ := ulid.New(ulid.Timestamp(now), ulid.Monotonic(rand.New(rand.NewSource(time.Unix(100000, 0).UnixNano())), 0))
	fmt.Printf("create: %s\n", u)
	return u.String()
}

func delete(ulid string) {
	fmt.Printf("delete: %s\n", ulid)
}
```
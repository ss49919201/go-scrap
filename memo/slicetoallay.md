```
s := make([]byte, 2, 4)
	s0 := (*[0]byte)(s) // s0 != nil
	s2 := (*[2]byte)(s) // &s2[0] == &s[0]

	fmt.Println(s0)     // &[]
	fmt.Println(&s2[0]) // &s[0]と同じアドレス
	fmt.Println(&s[0])  // &s2[0]と同じアドレス
```
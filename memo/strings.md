```
slice := []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh","xyzzy", "thud"}
joined := strings.Join(slice, " + ")
fmt.Println(joined)
// Output:
// foo + bar + baz + qux + quux + corge + grault + garply + waldo + fred + plugh + xyzzy + thud

splited := strings.Split(joined, " + ")
fmt.Println(splited)
// Output:
// [foo bar baz qux quux corge grault garply waldo fred plugh xyzzy thud]
```
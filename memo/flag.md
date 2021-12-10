# Parse

`*falg.FlagSet`型のメソッドの`Parse(arguments []string) error`を呼び出しています。
引数には os.Args[1:] が渡されていることから、コマンドライン引数をごにょごにょするのだろうと予測できます。
Parse メソッドの中身を見ていきます。

パッと見た感じだと、プライベートなフィールドの値を更新したり、受け取ったコマンドライン引数を受け取って処理しているようです。
詳しく見ていきます。

```go
	f.parsed = true
	f.args = arguments
```

二つのフィールドの値を更新しています。
`Parse(arguments []string) error`呼び出し済みフラグを立てています。
このフラグの値は、`Parsed()`を呼び出すことで得ることができます。

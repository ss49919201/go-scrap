# ビルド

## Build Constrains について

- ファイルがビルド時にパッケージの一部として含まれることに制約を付けられる
    - https://pkg.go.dev/cmd/go#hdr-Build_constraints

## mainパッケージに複数ファイルがある場合

- `go run *.go`
- `go build *.go`
- `go run hoge.go fuga.go`
- `go build hoge.go fuga.go`
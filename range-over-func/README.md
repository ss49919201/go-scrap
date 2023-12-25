```sh
go install golang.org/dl/gotip@latest
gotip download 510541
GOEXPERIMENT=range gotip run main.go
```

```
go install golang.org/dl/go1.22rc1@latest
go1.22rc1 download
GOEXPERIMENT=rangefunc go1.22rc1 run main.go
```
package main

import (
	"path/filepath"
	"runtime"
)

func main() {
	dir := allButLastElementOfPath(fileMain())
	println(dir)
	println(dir)
	println(dir)
}

func fileMain() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

func allButLastElementOfPath(path string) string {
	return filepath.Dir(path)
}

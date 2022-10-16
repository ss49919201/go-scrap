package main

import "fmt"

// interfaceAを満たしていないclassAを、interfaceが求められているケースで使用する際に、
// classAを改修せずに、interfaceAを満たすように実装したclassBを作成する。
// このclassBをAdapterと呼ぶ。
// classBにはclassAを委譲する。(継承するパターンもあるがGoではできない)
// 要はラッパークラスのようなもの。

type Cache struct{}

func (c *Cache) Set(v any) {
	fmt.Println("Set", v)
}

func (c *Cache) Get() {
	fmt.Println("Get")
}

type Loader interface {
	Load()
	Store(v any)
}

// Loader を満たすように実装したCacheAdapter
type CacheAdapter struct {
	*Cache
}

func (c *CacheAdapter) Load() {
	c.Cache.Get()
}

func (c *CacheAdapter) Store(v any) {
	c.Cache.Set(v)
}

func main() {
	var loader Loader
	// loader = &Cache{} // これはコンパイルエラーになる
	loader = &CacheAdapter{&Cache{}} // 今回のケースだとCacheはnilでも問題ないが、一応&Cache{}を渡している
	loader.Load()                    // Get
	loader.Store("hoge")             // Set hoge
}

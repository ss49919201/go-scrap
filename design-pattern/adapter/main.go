package main

import "github.com/go-tips/design-pattern/adapter/internal"

// interfaceAを満たしていないclassAを、interfaceが求められているケースで使用する際に、
// classAを改修せずに、interfaceAを満たすように実装したclassBを作成する。
// このclassBをAdapterと呼ぶ。
// classBにはclassAを委譲する。(継承するパターンもあるがGoではできない)
// 要はラッパークラスのようなもの。

func main() {
	var loader internal.Loader
	// loader = &Cache{} // これはコンパイルエラーになる
	loader = &internal.CacheAdapter{&internal.Cache{}} // 今回のケースだとCacheはnilでも問題ないが、一応&Cache{}を渡している
	loader.Load()                                      // Get
	loader.Store("hoge")                               // Set hoge
}

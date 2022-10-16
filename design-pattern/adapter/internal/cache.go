package internal

import "fmt"

// CacheはLoaderを実装していない
type Cache struct{}

func (c *Cache) Set(v any) {
	fmt.Println("Set", v)
}

func (c *Cache) Get() {
	fmt.Println("Get")
}

package internal

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

package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader/v7"
	lru "github.com/hashicorp/golang-lru"
)

func init() {
	c, _ := lru.NewARC(100)
	cache = &userCache[int, *User]{ARCCache: c}
}

var (
	cache *userCache[int, *User]
)

type userCache[K comparable, V any] struct {
	*lru.ARCCache
}

// Get gets an item from the cache
func (c *userCache[K, V]) Get(_ context.Context, key K) (dataloader.Thunk[V], bool) {
	v, ok := c.ARCCache.Get(key)
	if ok {
		return v.(dataloader.Thunk[V]), ok
	}
	return nil, ok
}

// Set sets an item in the cache
func (c *userCache[K, V]) Set(_ context.Context, key K, value dataloader.Thunk[V]) {
	c.ARCCache.Add(key, value)
}

// Delete deletes an item in the cache
func (c *userCache[K, V]) Delete(_ context.Context, key K) bool {
	if c.ARCCache.Contains(key) {
		c.ARCCache.Remove(key)
		return true
	}
	return false
}

// Clear cleasrs the cache
func (c *userCache[K, V]) Clear() {
	c.ARCCache.Purge()
}

var userData = map[int]*User{
	1: {1},
	2: {2},
	3: {3},
	4: {4},
}

type User struct {
	ID int `json:"id"`
}

type UserRepo struct{}

func (u *UserRepo) Find(id int) (*User, error) {
	user, ok := userData[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func main() {
	fn := func(ctx context.Context, keys []int) []*dataloader.Result[*User] {
		results := make([]*dataloader.Result[*User], len(keys))
		var wg sync.WaitGroup
		for i, key := range keys {
			wg.Add(1)
			go func(i, key int) {
				defer wg.Done()
				user, err := (&UserRepo{}).Find(key)
				results[i] = &dataloader.Result[*User]{Data: user, Error: err}
			}(i, key)
		}
		wg.Wait()
		return results
	}
	loader := dataloader.NewBatchedLoader(fn, dataloader.WithCache[int, *User](cache))

	loader.Load(context.Background(), 1)()
}

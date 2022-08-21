package main

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
	lru "github.com/hashicorp/golang-lru"
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

func NewUserCache[K comparable, V any](opts ...UserCacheOption[K, V]) *userCache[K, V] {
	u := &userCache[K, V]{}

	for _, opt := range opts {
		opt(u)
	}

	if u.ARCCache == nil {
		c, _ := lru.NewARC(100)
		u.ARCCache = c
	}

	return u
}

type UserCacheOption[K comparable, V any] func(*userCache[K, V])

func WithUserCache[K comparable, V any](c *lru.ARCCache) UserCacheOption[K, V] {
	return func(u *userCache[K, V]) {
		u.ARCCache = c
	}
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
	fmt.Println("UserRepo")
	user, ok := userData[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

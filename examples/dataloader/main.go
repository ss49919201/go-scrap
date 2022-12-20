package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	lru "github.com/hashicorp/golang-lru"
)

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

	c, _ := lru.NewARC(100)
	cache := NewUserCache(WithUserCache[int, *User](c))
	loader := dataloader.NewBatchedLoader(
		fn,
		dataloader.WithCache[int, *User](cache),
		dataloader.WithWait[int, *User](time.Second*2),
	)

	thunk := loader.Load(context.Background(), 1)

	// この時点でキャッシュにいるはず
	cachedThunk, _ := cache.Get(nil, 1)

	fmt.Println(thunk())

	// キャッシュされたクロージャーは更新されるのか？
	// インメモリはいけてた
	fmt.Println(cachedThunk())
}

package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
)

func runWithRedis(fn func()) {
	// poolを作成
	pool, err := newPool()
	if err != nil {
		panic(err)
	}

	// コンテナ起動
	resource, err := runRedis(pool)
	if err != nil {
		panic(err)
	}

	// リクエストを捌く準備ができるまで待つ
	if err = pool.Retry(func() error {
		return newRedisClient(resource.GetHostPort("6379/tcp")).
			Ping(context.Background()).
			Err()
	}); err != nil {
		panic(err)
	}

	// 何かしらテスト
	fn()

	// 掃除
	purgeRedis(pool, resource)
}

func newPool() (*dockertest.Pool, error) {
	return dockertest.NewPool("")
}

func runRedis(pool *dockertest.Pool) (*dockertest.Resource, error) {
	return pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "redis",
		Tag:        "latest",
	})
}

func closeRedis(resource *dockertest.Resource) {
	// エラーは無視
	resource.Close()
}

func purgeRedis(pool *dockertest.Pool, resource *dockertest.Resource) {
	// エラーは無視
	pool.Purge(resource)
}

func execRedis(resource *dockertest.Resource, cmd ...string) error {
	_, err := resource.Exec(cmd, dockertest.ExecOptions{})
	return err
}

func newRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func closeRedisClient(client *redis.Client) {
	// エラーは無視
	client.Close()
}

func fileMain() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

func allButLastElementOfPath(path string) string {
	return filepath.Dir(path)
}

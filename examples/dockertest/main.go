package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
)

func main() {
	// creates a docker client
	pool, err := newPool()
	if err != nil {
		panic(err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := runRedis(pool)
	if err != nil {
		panic(err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		return newRedisClient(resource.GetHostPort("6379/tcp")).Ping(context.Background()).Err()
	}); err != nil {
		panic(err)
	}

	// You can't defer this because os.Exit doesn't care for defer
	if err = pool.Purge(resource); err != nil {
		panic(err)
	}
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

package cache_test

import (
	"fmt"
	"testing"

	"github.com/gomig/cache"
	"github.com/redis/go-redis/v9"
)

func redisQueue() cache.Queue {
	return cache.NewRedisQueue("test", redis.Options{Addr: "localhost:6379"})
}

func TestRedisQueue(t *testing.T) {
	if err := redisQueue().Push("john"); err != nil {
		t.Fatal(err)
	}

	if err := redisQueue().Push("jack"); err != nil {
		t.Fatal(err)
	}

	if v, err := redisQueue().Pull(); err != nil {
		t.Fatal(err)
	} else if v == nil {
		fmt.Println("nil result")
	} else {
		fmt.Println("result: ", *v)
	}
}

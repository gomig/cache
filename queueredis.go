package cache

import (
	"context"
	"errors"

	"github.com/gomig/utils"
	"github.com/redis/go-redis/v9"
)

type rQueue struct {
	name   string
	client *redis.Client
}

func (rQueue) err(pattern string, params ...any) error {
	return utils.TaggedError([]string{"RedisQueue"}, pattern, params...)
}

func (rq *rQueue) init(name string, opt redis.Options) {
	rq.name = name
	rq.client = redis.NewClient(&opt)
}

func (rq rQueue) Push(value any) error {
	if err := rq.client.LPush(context.TODO(), rq.name, value).Err(); err != nil {
		return rq.err(err.Error())
	}
	return nil
}

func (rq rQueue) Pull() (*string, error) {
	v, err := rq.client.RPop(context.TODO(), rq.name).Result()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, rq.err(err.Error())
	} else if v == "" {
		return nil, nil
	} else {
		return &v, nil
	}
}

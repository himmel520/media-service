package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New(conn string) (*redis.Client, error) {
	opt, err := redis.ParseURL(conn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(context.Background()).Result()

	return rdb, err
}

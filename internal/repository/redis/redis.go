package redis

import (
	"context"
	"time"

	"github.com/himmel520/uoffer/mediaAd/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb        *redis.Client
	expiration time.Duration
}

func New(cfg *config.Cache) (*Redis, error) {
	opt, err := redis.ParseURL(cfg.Conn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(context.Background()).Result()
	return &Redis{rdb: rdb, expiration: cfg.Exp}, err
}

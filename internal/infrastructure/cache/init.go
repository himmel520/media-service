package cache

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

//go:generate mockery --all

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, prefix string) error
}

func Init(rdb *goredis.Client, exp time.Duration) Cache {
	return NewCache(rdb, exp)
}

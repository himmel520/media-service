package cache

import (
	"context"
	"time"

	"github.com/himmel520/media-service/internal/infrastructure/cache/redis"
	goredis "github.com/redis/go-redis/v9"
)

//go:generate mockery --all

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, prefix string) error
}

func New(rdb *goredis.Client, exp time.Duration) Cache {
	return redis.NewCache(rdb, exp)
}

package cache

import (
	"context"
	"time"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/redis"
	goredis "github.com/redis/go-redis/v9"
)

type (
	Cache struct {
		Adv AdvCache
	}

	AdvCache interface {
		Set(ctx context.Context, key string, advs []*entity.AdvResponse) error
		Get(ctx context.Context, key string) ([]*entity.AdvResponse, error)
		Delete(ctx context.Context) error
	}
)

func New(rdb *goredis.Client, exp time.Duration) *Cache {
	return &Cache{
		Adv: redis.NewAdvCache(rdb, exp),
	}
}

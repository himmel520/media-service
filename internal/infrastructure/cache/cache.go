package cache

import (
	"time"

	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/redis"
	goredis "github.com/redis/go-redis/v9"
)

//go:generate mockery --all
type Cache struct {
	client *redis.Cache
}

func New(db *goredis.Client, exp time.Duration) *Cache {
	client := redis.NewClient(db, exp)
	return &Cache{client: client}
}

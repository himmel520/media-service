package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/errcache"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
	exp time.Duration
}

func New(conn string) (*redis.Client, error) {
	opt, err := redis.ParseURL(conn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(context.Background()).Result()

	return rdb, err
}

func NewClient(db *redis.Client, exp time.Duration) *Cache {
	return &Cache{rdb: db, exp: exp}

}

func (r *Cache) Set(ctx context.Context, key string, advs any) error {
	advsByte, err := json.Marshal(advs)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, key, string(advsByte), r.exp).Result()
	return err
}

func (r *Cache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errcache.ErrKeyNotFound
		}
		return "", err
	}

	return val, err
}

func (r *Cache) Delete(ctx context.Context) error {
	keys, err := r.rdb.Keys(ctx, "advs:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		err := r.rdb.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

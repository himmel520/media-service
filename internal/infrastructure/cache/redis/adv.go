package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/errcache"

	"github.com/redis/go-redis/v9"
)

type AdvCache struct {
	rdb *redis.Client
	exp time.Duration
}

func NewAdvCache(rdb *redis.Client, exp time.Duration) *AdvCache {
	return &AdvCache{rdb: rdb, exp: exp}
}

func (r *AdvCache) Set(ctx context.Context, key string, advs []*entity.AdvResponse) error {
	advsByte, err := json.Marshal(advs)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, key, string(advsByte), r.exp).Result()
	return err
}

func (r *AdvCache) Get(ctx context.Context, key string) ([]*entity.AdvResponse, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errcache.ErrKeyNotFound
		}
		return nil, err
	}

	advs := []*entity.AdvResponse{}
	err = json.Unmarshal([]byte(val), &advs)

	return advs, err
}

func (r *AdvCache) Delete(ctx context.Context) error {
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

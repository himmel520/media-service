package redis

import (
	"context"
	"encoding/json"

	"github.com/himmel520/uoffer/mediaAd/internal/repository"
	"github.com/himmel520/uoffer/mediaAd/models"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) SetAdv(ctx context.Context, key string, advs []*models.AdvResponse) error {
	advsByte, err := json.Marshal(advs)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, key, string(advsByte), r.expiration).Result()
	return err
}

func (r *Redis) GetAdv(ctx context.Context, key string) ([]*models.AdvResponse, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, repository.ErrKeyNotFound
		}
		return nil, err
	}

	advs := []*models.AdvResponse{}
	err = json.Unmarshal([]byte(val), &advs)

	return advs, err
}

func (r *Redis) DeleteAdvsCache(ctx context.Context) error {
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

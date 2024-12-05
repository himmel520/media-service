package cache

import (
	"context"
	"encoding/json"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Cache keys
const (
	AdvPrefixKey = "advs:*"

	LogoPrefixKey = "logo:*"
	AllLogoskey   = "logo:all"
)

func NewRedis(conn string) (*goredis.Client, error) {
	opt, err := goredis.ParseURL(conn)
	if err != nil {
		return nil, err
	}

	rdb := goredis.NewClient(opt)

	_, err = rdb.Ping(context.Background()).Result()

	return rdb, err
}

type Redis struct {
	rdb *goredis.Client
	exp time.Duration
}

func NewCache(db *goredis.Client, exp time.Duration) *Redis {
	return &Redis{rdb: db, exp: exp}
}

func (r *Redis) Set(ctx context.Context, key string, value any) error {
	byte, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, key, string(byte), r.exp).Result()
	return err
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == goredis.Nil {
			return "", ErrKeyNotFound
		}
		return "", err
	}

	return val, err
}

func (r *Redis) Delete(ctx context.Context, prefix string) error {
	keys, err := r.rdb.Keys(ctx, prefix).Result()
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

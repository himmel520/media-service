package adUC

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache/errcache"
)

func (uc *AdUC) GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error) {
	key := generateCacheKeyAdv(limit, offset, posts, priority)

	advs := []*entity.AdvResponse{}
	val, err := uc.cache.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, errcache.ErrKeyNotFound) {
			uc.log.Error(err)
		}

		advs, err = uc.repo.GetAllWithFilter(ctx, limit, offset, posts, priority)
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(context.Background(), key, advs); err != nil {
			uc.log.Error(err)
		}

		return advs, nil
	}

	err = json.Unmarshal([]byte(val), &advs)
	return advs, err
}

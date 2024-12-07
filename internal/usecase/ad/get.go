package adUC

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/convert"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
	log "github.com/youroffer/logger"
)

func generateCacheKeyAdv(posts []string, priority []int) string {
	key := fmt.Sprintf("%s:%s",
		strings.Join(posts, ","),
		strings.Join(convert.ApplyToSlice(priority, func(i int) string { return string(i) }), ","))

	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return cache.AdvPrefixKey + hash
}

func (uc *AdUC) Get(ctx context.Context, params usecase.PageParams) (*entity.AdsResp, error) {
	ads, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.AdsResp{
		Data:    ads,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

func (uc *AdUC) GetWithFilter(ctx context.Context, params usecase.AdvFilterParams) ([]*entity.AdvResp, error) {
	var ads []*entity.AdvResp

	cacheKey := generateCacheKeyAdv(params.Posts, params.Priority)

	bytes, err := uc.cache.Get(ctx, cacheKey)
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
			log.ErrMsg(err, "get ads cache")
		}

		ads, err := uc.repo.GetWithFilter(ctx, uc.db.DB(), repository.AdvFilterParams{
			Posts:    params.Posts,
			Priority: params.Priority})
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(ctx, cacheKey, ads); err != nil {
			log.ErrMsg(err, "set ads cache")
		}

		return ads, nil
	}

	err = json.Unmarshal([]byte(bytes), &ads)
	return ads, err
}

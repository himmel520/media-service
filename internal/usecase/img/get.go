package imgUC

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
	log "github.com/youroffer/logger"
)

func (uc *ImgUC) Get(ctx context.Context, params usecase.PageParams) (*entity.ImagesResp, error) {
	images, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.ImagesResp{
		Data:    images,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

func (uc *ImgUC) GetAllLogos(ctx context.Context) (entity.LogosResp, error) {
	var logos entity.LogosResp
	
	bytes, err := uc.cache.Get(ctx, cache.AllLogoskey)
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
			log.Err(err)
		}
		
		logos, err := uc.repo.GetAllLogos(ctx, uc.db.DB())
		if err != nil {
			return nil, err
		}
		
		if err = uc.cache.Set(ctx, cache.AllLogoskey, logos); err != nil {
			log.ErrMsg(err, "set all logos cache")
		}
		
		return logos, nil
	}

	err = json.Unmarshal([]byte(bytes), &logos)
	return logos, err
}

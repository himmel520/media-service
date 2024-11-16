package imgUC

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
)

func (uc *ImgUC) GetByID(ctx context.Context, id int) (*entity.Image, error) {
	// return uc.repo.GetByID(ctx, id)
	return nil, nil
}

func (uc *ImgUC) Get(ctx context.Context, params usecase.PageParams) (*entity.ImagesResp, error) {
	images, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, err
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

	bytes, err := uc.cache.Get(ctx, allLogosCachekey)
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
			uc.log.Error(err)
		}

		logos, err := uc.repo.GetAllLogos(ctx, uc.db.DB())
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(ctx, allLogosCachekey, logos); err != nil {
			uc.log.Error(err)
		}

		return logos, nil
	}

	err = json.Unmarshal([]byte(bytes), &logos)
	return logos, err
}

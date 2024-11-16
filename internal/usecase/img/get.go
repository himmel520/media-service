package imgUC

import (
	"context"
	"encoding/json"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
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
	bytes, err := uc.cache.GetAll(ctx, logoCachePrefix)
	if err != nil {
		return nil, err

	}
	res := make(entity.LogosResp, len(bytes))
	for _, val := range bytes {
		var logo api.LogosRespItem
		err := json.Unmarshal([]byte(val), &logo)
		if err != nil {
			return nil, err
		}
		res[logo.ID] = logo
	}

	return res
	// return uc.repo.GetAllLogos(ctx, uc.db.DB())
}

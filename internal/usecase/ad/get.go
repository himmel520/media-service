package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
)

func (uc *AdUC) Get(ctx context.Context, params usecase.PageParams) (*entity.AdsResp, error) {
	ads, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, err
	}

	return &entity.AdsResp{
		Data:    ads,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

func (uc *AdUC) GetWithFilter(ctx context.Context, params usecase.AdvFilterParams) ([]*entity.AdvResp, error) {
	return uc.repo.GetWithFilter(ctx, uc.db.DB(), repository.AdvFilterParams{
		Posts: params.Posts,
		Priority: params.Priority,
	})
}

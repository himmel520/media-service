package colorUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
)

func (uc *ColorUC) Get(ctx context.Context, params usecase.PageParams) (*entity.ColorsResp, error) {
	colors, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, err
	}

	return &entity.ColorsResp{
		Data:    colors,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

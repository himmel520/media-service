package tgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/lib/paging"
	"github.com/himmel520/media-service/internal/usecase"
)

func (uc *TgUC) Get(ctx context.Context, params usecase.PageParams) (*entity.TgsResp, error) {
	tgs, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.TgsResp{
		Data:    tgs,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *TgUC) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.TGsResp, error) {
	tgs, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.TGsResp{
		TGs:   tgs,
		Total: count,
	}, err
}

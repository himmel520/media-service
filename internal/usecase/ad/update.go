package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *AdUC) Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResponse, error) {
	if err := uc.repo.Update(ctx, id, adv); err != nil {
		return nil, err
	}

	return uc.repo.GetByID(ctx, id)
}

package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *AdUC) Add(ctx context.Context, adv *entity.Adv) (*entity.AdvResponse, error) {
	id, err := uc.repo.Add(ctx, adv)
	if err != nil {
		return nil, err
	}

	return uc.repo.GetByID(ctx, id)
}

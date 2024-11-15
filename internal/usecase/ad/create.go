package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *AdUC) Create(ctx context.Context, adv *entity.Adv) (*entity.AdvResp, error) {
	return uc.repo.Create(ctx, uc.db.DB(), adv)
}

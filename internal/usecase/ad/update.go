package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *AdUC) Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResp, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, adv)
}

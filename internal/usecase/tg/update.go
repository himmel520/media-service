package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *TgUC) Update(ctx context.Context, id int, tg *entity.TgUpdate) (*entity.Tg, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, tg)
}

package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *TgUC) Create(ctx context.Context, tg *entity.Tg) (*entity.Tg, error) {
	return uc.repo.Create(ctx, uc.db.DB(), tg)
}

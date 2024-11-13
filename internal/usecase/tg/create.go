package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *TgUC) Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error) {
	return uc.repo.Add(ctx, tg)
}

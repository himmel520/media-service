package colorUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ColorUC) Create(ctx context.Context, color *entity.Color) (*entity.Color, error) {
	return uc.repo.Create(ctx, uc.db.DB(), color)
}

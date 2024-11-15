package colorUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ColorUC) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.Color, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, color)
}

package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ImgUC) Create(ctx context.Context, image *entity.Image) (*entity.Image, error) {
	return uc.repo.Create(ctx, uc.db.DB(), image)
}

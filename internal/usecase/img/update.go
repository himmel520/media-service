package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ImgUC) Update(ctx context.Context, id int, image *entity.ImageUpdate) (*entity.Image, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, image)
}

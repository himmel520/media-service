package imgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ImgUC) Update(ctx context.Context, id int, image *entity.ImageUpdate) (*entity.Image, error) {
	newImage, err := uc.repo.Update(ctx, uc.db.DB(), id, image)
	if err != nil {
		return nil, fmt.Errorf("update image: %w", err)
	}

	uc.DeleteImageCache(ctx, newImage.Type)

	return newImage, nil
}

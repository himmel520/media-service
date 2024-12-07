package imgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ImgUC) Create(ctx context.Context, image *entity.Image) (*entity.Image, error) {
	newImage, err := uc.repo.Create(ctx, uc.db.DB(), image)
	if err != nil {
		return nil, fmt.Errorf("create img: %w", err)
	}

	uc.DeleteImageCache(ctx, newImage.Type)

	return newImage, nil
}

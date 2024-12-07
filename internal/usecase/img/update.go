package imgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *ImgUC) Update(ctx context.Context, id int, image *entity.ImageUpdate) (*entity.Image, error) {
	newImage, err := uc.repo.Update(ctx, uc.db.DB(), id, image)
	if err != nil {
		return nil, fmt.Errorf("update image: %w", err)
	}

	switch newImage.Type {
	case entity.ImageTypeLogo:
		err = uc.cache.Delete(context.Background(), cache.LogoPrefixKey)
	case entity.ImageTypeAdv:
		err = uc.cache.Delete(context.Background(), cache.AdvPrefixKey)
	}

	if err != nil {
		log.ErrMsgf(err, "delete %s cache", newImage.Type)
	}

	return newImage, nil
}

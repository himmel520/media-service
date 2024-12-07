package imgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *ImgUC) Create(ctx context.Context, image *entity.Image) (*entity.Image, error) {
	newImage, err := uc.repo.Create(ctx, uc.db.DB(), image)
	if err != nil {
		return nil, fmt.Errorf("create img: %w", err)
	}

	if err := uc.cache.Delete(context.Background(), cache.LogoPrefixKey); err != nil {
		log.ErrMsg(err, "delete logo cache")
	}

	return newImage, nil
}

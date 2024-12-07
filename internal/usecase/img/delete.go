package imgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *ImgUC) Delete(ctx context.Context, id int) error {
	imageType, err := uc.repo.GetImageTypeByID(ctx, uc.db.DB(), id)
	if err != nil {
		return fmt.Errorf("get image type: %w", err)
	}

	if err := uc.repo.Delete(ctx, uc.db.DB(), id); err != nil {
		return fmt.Errorf("delete img: %w", err)
	}

	if imageType == entity.ImageTypeLogo {
		if err := uc.cache.Delete(context.Background(), cache.LogoPrefixKey); err != nil {
			log.ErrMsg(err, "delete logo cache")
		}
	}

	return nil
}

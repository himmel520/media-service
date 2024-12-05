package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *ImgUC) DeleteCache(ctx context.Context) {
	if err := uc.cache.Delete(context.Background(), cache.LogoPrefixKey); err != nil {
		log.Err(err)
	}
}

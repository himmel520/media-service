package imgUC

import (
	"context"

	log "github.com/youroffer/logger"
)

const (
	logoCachePrefix  = "logo:*"
	allLogosCachekey = "logo:all"
)

func (uc *ImgUC) DeleteCache(ctx context.Context) {
	if err := uc.cache.Delete(context.Background(), logoCachePrefix); err != nil {
		log.Err(err)
	}
}

package adUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *AdUC) Delete(ctx context.Context, id int) error {
	if err := uc.repo.Delete(ctx, uc.db.DB(), id); err != nil {
		return fmt.Errorf("delete adv: %w", err)
	}

	if err := uc.cache.Delete(ctx, cache.AdvPrefixKey); err != nil {
		log.ErrMsg(err, "delete ads cache")
	}

	return nil
}

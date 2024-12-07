package adUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *AdUC) Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResp, error) {
	updatedAdv, err := uc.repo.Update(ctx, uc.db.DB(), id, adv)
	if err != nil {
		return nil, fmt.Errorf("update adv: %w", err)
	}

	if err := uc.cache.Delete(ctx, cache.AdvPrefixKey); err != nil {
		log.ErrMsg(err, "delete ads cache")
	}

	return updatedAdv, nil
}

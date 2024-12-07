package adUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *AdUC) Create(ctx context.Context, adv *entity.Adv) (*entity.AdvResp, error) {
	newAdv, err := uc.repo.Create(ctx, uc.db.DB(), adv)
	if err != nil {
		return nil, fmt.Errorf("create adv: %w", err)
	}

	if err = uc.cache.Delete(ctx, cache.AdvPrefixKey); err != nil {
		log.ErrMsg(err, "delete ads cache")
	}
	
	return newAdv, nil
}

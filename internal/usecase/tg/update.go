package tgUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *TgUC) Update(ctx context.Context, id int, tg *entity.TgUpdate) (*entity.Tg, error) {
	updatedTg, err := uc.repo.Update(ctx, uc.db.DB(), id, tg)
	if err != nil {
		return nil, fmt.Errorf("update tg: %w", err)
	}

	if err := uc.cache.Delete(ctx, cache.AdvPrefixKey); err != nil {
		log.ErrMsg(err, "delete ads cache")
	}

	return updatedTg, nil
}

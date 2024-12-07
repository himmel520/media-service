package colorUC

import (
	"context"
	"fmt"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	log "github.com/youroffer/logger"
)

func (uc *ColorUC) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.Color, error) {
	updatedColor, err := uc.repo.Update(ctx, uc.db.DB(), id, color)
	if err != nil {
		return nil, fmt.Errorf("update color: %w", err)
	}

	if err := uc.cache.Delete(ctx, cache.AdvPrefixKey); err != nil {
		log.ErrMsg(err, "delete ads cache")
	}

	return updatedColor, nil
}

package color

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminColorsPost(ctx context.Context, req *api.ColorPost) (api.V1AdminColorsPostRes, error) {
	color, err := h.uc.Create(ctx, &entity.Color{
		Title: req.GetTitle(),
		Hex:   req.GetHex(),
	})

	switch {
	case errors.Is(err, repoerr.ErrColorHexExist):
		return &api.V1AdminColorsPostConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err)
		return nil, err
	}
	
	return entity.ColorToApi(color), nil
}

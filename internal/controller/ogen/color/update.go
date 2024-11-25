package color

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminColorsIDPut(ctx context.Context, req *api.ColorPut, params api.V1AdminColorsIDPutParams) (api.V1AdminColorsIDPutRes, error) {
	newColor := &entity.ColorUpdate{
		Title: entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Hex:   entity.Optional[string]{Value: req.GetHex().Value, Set: req.GetHex().Set},
	}

	if !newColor.IsSet() {
		return &api.V1AdminColorsIDPutBadRequest{Message: "no changes"}, nil
	}

	color, err := h.uc.Update(ctx, params.ID, newColor)
	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		return &api.V1AdminColorsIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrColorHexExist):
		return &api.V1AdminColorsIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, map[string]string{
			"req_id": middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.ColorToApi(color), nil
}

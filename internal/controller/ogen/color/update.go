package color

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminColorsIDPut(ctx context.Context, req *api.ColorPut, params api.V1AdminColorsIDPutParams) (api.V1AdminColorsIDPutRes, error) {
	newColor := &entity.ColorUpdate{
		Title: entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Hex:   entity.Optional[string]{Value: req.GetHex().Value, Set: req.GetHex().Set},
	}

	if !newColor.IsSet() {
		return &api.V1AdminColorsIDPutBadRequest{Message: "no changes"}, nil
	}

	Color, err := h.uc.Update(ctx, params.ID, newColor)
	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		return &api.V1AdminColorsIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrColorHexExist):
		return &api.V1AdminColorsIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return entity.ColorToApi(Color), nil
}

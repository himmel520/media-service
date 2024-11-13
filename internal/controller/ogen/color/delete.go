package color

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminColorsIDDelete(ctx context.Context, params api.V1AdminColorsIDDeleteParams) (api.V1AdminColorsIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)
	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		return &api.V1AdminColorsIDDeleteNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrColorDependencyExist):
		return &api.V1AdminColorsIDDeleteConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return &api.V1AdminColorsIDDeleteNoContent{}, nil
}

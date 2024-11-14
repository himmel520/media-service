package ad

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminAdsIDDelete(ctx context.Context, params api.V1AdminAdsIDDeleteParams) (api.V1AdminAdsIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)
	
	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		return &api.V1AdminAdsIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return &api.V1AdminAdsIDDeleteNoContent{}, nil
}

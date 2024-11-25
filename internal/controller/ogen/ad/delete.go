package ad

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminAdsIDDelete(ctx context.Context, params api.V1AdminAdsIDDeleteParams) (api.V1AdminAdsIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)

	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		return &api.V1AdminAdsIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, map[string]interface{}{
			"req_id": middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return &api.V1AdminAdsIDDeleteNoContent{}, nil
}

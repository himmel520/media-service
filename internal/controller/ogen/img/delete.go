package img

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminImagesIDDelete(ctx context.Context, params api.V1AdminImagesIDDeleteParams) (api.V1AdminImagesIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)
	switch {
	case errors.Is(err, repoerr.ErrImageNotFound):
		return &api.V1AdminImagesIDDeleteNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrImageDependencyExist):
		return &api.V1AdminImagesIDDeleteConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return &api.V1AdminImagesIDDeleteNoContent{}, nil
}

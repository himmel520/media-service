package tg

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminTgsIDDelete(ctx context.Context, params api.V1AdminTgsIDDeleteParams) (api.V1AdminTgsIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)
	switch {
	case errors.Is(err, repoerr.ErrTGNotFound):
		return &api.V1AdminTgsIDDeleteNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrTGDependencyExist):
		return &api.V1AdminTgsIDDeleteConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err)
		return nil, err
	}

	return &api.V1AdminTgsIDDeleteNoContent{}, nil
}

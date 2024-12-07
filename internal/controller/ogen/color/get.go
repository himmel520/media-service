package color

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/controller/ogen"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/internal/usecase"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminColorsGet(ctx context.Context, params api.V1AdminColorsGetParams) (api.V1AdminColorsGetRes, error) {
	colorsResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		return &api.V1AdminColorsGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return colorsResp.ToApi(), nil
}

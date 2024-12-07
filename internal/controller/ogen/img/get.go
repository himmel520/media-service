package img

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/controller/ogen"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/internal/usecase"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1LogosGet(ctx context.Context) (api.V1LogosGetRes, error) {
	logos, err := h.uc.GetAllLogos(ctx)

	switch {
	case errors.Is(err, repoerr.ErrImageNotFound):
		return &api.V1LogosGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.LogosRespToApi(logos), nil
}

func (h *Handler) V1AdminImagesGet(ctx context.Context, params api.V1AdminImagesGetParams) (api.V1AdminImagesGetRes, error) {
	imagesResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrImageNotFound):
		return &api.V1AdminImagesGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return imagesResp.ToApi(), nil
}

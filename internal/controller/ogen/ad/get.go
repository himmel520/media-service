package ad

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/controller/ogen"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/internal/lib/convert"
	"github.com/himmel520/media-service/internal/usecase"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdsGet(ctx context.Context, params api.V1AdsGetParams) (api.V1AdsGetRes, error) {
	ads, err := h.uc.GetWithFilter(ctx, usecase.AdvFilterParams{
		Posts: params.Post,
		Priority: convert.ApplyToSlice(
			params.Priority,
			func(item api.V1AdsGetPriorityItem) int {
				return int(item)
			}),
	})

	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		return &api.V1AdsGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err)
		return nil, err
	}

	return entity.AdsToApi(ads), nil
}

func (h *Handler) V1AdminAdsGet(ctx context.Context, params api.V1AdminAdsGetParams) (api.V1AdminAdsGetRes, error) {
	adsResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		return &api.V1AdminAdsGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err)
		return nil, err
	}

	return adsResp.ToApi(), nil
}

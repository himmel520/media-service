package tg

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/controller/ogen"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/internal/usecase"
)

func (h *Handler) V1AdminTgsGet(ctx context.Context, params api.V1AdminTgsGetParams) (api.V1AdminTgsGetRes, error) {
	tgsResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrTGNotFound):
		return &api.V1AdminTgsGetNotFound{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return tgsResp.ToApi(), nil
}

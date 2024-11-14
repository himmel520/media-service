package tg

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminTgsPost(ctx context.Context, req *api.TgPost) (api.V1AdminTgsPostRes, error) {
	tg, err := h.uc.Create(ctx, &entity.Tg{
		Url:   req.URL.String(),
		Title: req.GetTitle(),
	})

	switch {
	case errors.Is(err, repoerr.ErrTGExist):
		return &api.V1AdminTgsPostConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return entity.TgToApi(tg), nil
}

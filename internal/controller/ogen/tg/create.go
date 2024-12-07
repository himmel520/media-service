package tg

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
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
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.TgToApi(tg), nil
}

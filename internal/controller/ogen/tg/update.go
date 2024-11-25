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

func (h *Handler) V1AdminTgsIDPut(ctx context.Context, req *api.TgPut, params api.V1AdminTgsIDPutParams) (api.V1AdminTgsIDPutRes, error) {
	newTg := &entity.TgUpdate{
		Title: entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Url:   entity.Optional[string]{Value: req.URL.Value.String(), Set: req.GetURL().Set},
	}

	if !newTg.IsSet() {
		return &api.V1AdminTgsIDPutBadRequest{Message: "no changes"}, nil
	}

	tg, err := h.uc.Update(ctx, params.ID, newTg)
	switch {
	case errors.Is(err, repoerr.ErrTGNotFound):
		return &api.V1AdminTgsIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrTGExist):
		return &api.V1AdminTgsIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, map[string]string{
			"req_id": middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.TgToApi(tg), nil
}

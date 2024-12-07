package ad

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminAdsIDPut(ctx context.Context, req *api.AdPut, params api.V1AdminAdsIDPutParams) (api.V1AdminAdsIDPutRes, error) {
	newAd := &entity.AdvUpdate{
		ColorID:     entity.Optional[int]{Value: req.GetColorsID().Value, Set: req.GetColorsID().IsSet()},
		ImageID:     entity.Optional[int]{Value: req.GetImagesID().Value, Set: req.GetImagesID().IsSet()},
		TgID:        entity.Optional[int]{Value: req.GetTgID().Value, Set: req.GetTgID().IsSet()},
		Post:        entity.Optional[string]{Value: req.GetPost().Value, Set: req.GetPost().IsSet()},
		Title:       entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().IsSet()},
		Description: entity.Optional[string]{Value: req.GetDescription().Value, Set: req.GetDescription().IsSet()},
		Priority:    entity.Optional[int]{Value: int(req.GetPriority().Value), Set: req.GetPriority().IsSet()},
	}

	if !newAd.IsSet() {
		return &api.V1AdminAdsIDPutBadRequest{Message: "no changes"}, nil
	}

	ad, err := h.uc.Update(ctx, params.ID, newAd)
	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		return &api.V1AdminAdsIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrAdvDependencyNotExist):
		return &api.V1AdminAdsIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.AdvRespToApi(ad), nil
}

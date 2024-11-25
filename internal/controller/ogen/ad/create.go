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

func (h *Handler) V1AdminAdsPost(ctx context.Context, req *api.AdPost) (api.V1AdminAdsPostRes, error) {
	adv, err := h.uc.Create(ctx, &entity.Adv{
		ColorID:     req.GetColorsID(),
		ImageID:     req.GetImagesID(),
		TgID:        req.GetTgID(),
		Post:        req.GetPost(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Priority:    int(req.GetPriority()),
	})

	switch {
	case errors.Is(err, repoerr.ErrAdvDependencyNotExist):
		return &api.V1AdminAdsPostConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, map[string]interface{}{
			"req_id": middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.AdvRespToApi(adv), nil
}

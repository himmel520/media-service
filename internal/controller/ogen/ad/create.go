package ad

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/media-service/pkg/log"
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
		log.Error(err)
		return nil, err
	}

	return entity.AdvRespToApi(adv), nil
}

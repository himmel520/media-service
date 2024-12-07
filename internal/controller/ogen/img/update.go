package img

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminImagesIDPut(ctx context.Context, req *api.ImagePut, params api.V1AdminImagesIDPutParams) (api.V1AdminImagesIDPutRes, error) {
	newImage := &entity.ImageUpdate{
		Title: entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Url:   entity.Optional[string]{Value: req.URL.Value.String(), Set: req.GetURL().Set},
		Type:  entity.Optional[string]{Value: string(req.GetType().Value), Set: req.GetType().Set},
	}

	if !newImage.IsSet() {
		return &api.V1AdminImagesIDPutBadRequest{Message: "no changes"}, nil
	}

	image, err := h.uc.Update(ctx, params.ID, newImage)
	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		return &api.V1AdminImagesIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrColorHexExist):
		return &api.V1AdminImagesIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.ImageToApi(image), nil
}

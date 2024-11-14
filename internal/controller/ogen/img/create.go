package img

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminImagesPost(ctx context.Context, req *api.ImagePost) (api.V1AdminImagesPostRes, error) {
	image, err := h.uc.Create(ctx, &entity.Image{
		Url:   req.URL.String(),
		Title: req.GetTitle(),
		Type:  string(req.GetType()),
	})

	switch {
	case errors.Is(err, repoerr.ErrImageExist):
		return &api.V1AdminImagesPostConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}
	
	return entity.ImageToApi(image), nil
}

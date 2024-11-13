package img

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminImagesPost(ctx context.Context, req *api.ImagePost) (api.V1AdminImagesPostRes, error) {
	// Создание нового изображения
	return &api.Image{}, nil
}

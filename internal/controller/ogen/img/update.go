package img

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminImagesIDPut(ctx context.Context, req *api.ImagePut, params api.V1AdminImagesIDPutParams) (api.V1AdminImagesIDPutRes, error) {
	// Обновление данных изображения по ID
	return &api.Image{}, nil
}

package img

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminImagesIDDelete(ctx context.Context, params api.V1AdminImagesIDDeleteParams) (api.V1AdminImagesIDDeleteRes, error) {
	// Удаление изображения по ID
	return &api.V1AdminImagesIDDeleteNoContent{}, nil
}

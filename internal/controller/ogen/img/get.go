package img

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1AdminImagesGet(ctx context.Context, params api.V1AdminImagesGetParams) (api.V1AdminImagesGetRes, error) {
	// Получение списка изображений с пагинацией
	return &api.ImagesResp{}, nil
}

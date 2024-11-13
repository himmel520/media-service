package img

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1LogosGet(ctx context.Context) (api.V1LogosGetRes, error) {
	// Возвращает список всех логотипов
	return &api.LogosResp{}, nil
}

func (h *Handler) V1AdminImagesGet(ctx context.Context, params api.V1AdminImagesGetParams) (api.V1AdminImagesGetRes, error) {
	// Получение списка изображений с пагинацией
	return &api.ImagesResp{}, nil
}

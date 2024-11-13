package color

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1LogosGet(ctx context.Context) (api.V1LogosGetRes, error) {
	// Возвращает список всех логотипов
	return &api.LogosResp{}, nil
}

func (h *Handler) V1AdminColorsGet(ctx context.Context, params api.V1AdminColorsGetParams) (api.V1AdminColorsGetRes, error) {
	// Получение списка цветов с пагинацией
	return &api.ColorsResp{}, nil
}

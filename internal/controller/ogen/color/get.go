package color

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminColorsGet(ctx context.Context, params api.V1AdminColorsGetParams) (api.V1AdminColorsGetRes, error) {
	// Получение списка цветов с пагинацией
	return &api.ColorsResp{}, nil
}

package color

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminColorsIDDelete(ctx context.Context, params api.V1AdminColorsIDDeleteParams) (api.V1AdminColorsIDDeleteRes, error) {
	// Удаление цвета по ID
	return &api.V1AdminColorsIDDeleteNoContent{}, nil
}

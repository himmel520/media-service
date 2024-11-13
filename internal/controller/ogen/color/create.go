package color

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminColorsPost(ctx context.Context, req *api.ColorPost) (api.V1AdminColorsPostRes, error) {
	// Создание нового цвета
	return &api.Color{}, nil
}

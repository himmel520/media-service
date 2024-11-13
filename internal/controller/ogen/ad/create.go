package ad

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdminAdsPost(ctx context.Context, req *api.AdPost) (api.V1AdminAdsPostRes, error) {
	// Создание новой рекламы
	return &api.Ad{}, nil
}

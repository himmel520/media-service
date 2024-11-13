package ad

import (
	"context"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) V1AdsGet(ctx context.Context, params api.V1AdsGetParams) (api.V1AdsGetRes, error) {
	// Получение списка реклам с фильтрацией по приоритету и должности
	return &api.Ads{}, nil
}

func (h *Handler) V1AdminAdsGet(ctx context.Context, params api.V1AdminAdsGetParams) (api.V1AdminAdsGetRes, error) {
	// Получение списка всех реклам с пагинацией
	return &api.AdsResp{}, nil
}

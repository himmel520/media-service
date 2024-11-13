package ad

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1AdminAdsIDDelete(ctx context.Context, params api.V1AdminAdsIDDeleteParams) (api.V1AdminAdsIDDeleteRes, error) {
	// Удаление рекламы по ID
	return &api.V1AdminAdsIDDeleteNoContent{}, nil
}

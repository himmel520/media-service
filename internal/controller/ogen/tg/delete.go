package tg

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1AdminTgsIDDelete(ctx context.Context, params api.V1AdminTgsIDDeleteParams) (api.V1AdminTgsIDDeleteRes, error) {
	// Удаление тг по ID
	return &api.V1AdminTgsIDDeleteNoContent{}, nil
}

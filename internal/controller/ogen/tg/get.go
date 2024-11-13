package tg

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1AdminTgsGet(ctx context.Context, params api.V1AdminTgsGetParams) (api.V1AdminTgsGetRes, error) {
	// Получение списка тг с пагинацией
	return &api.TgsResp{}, nil
}
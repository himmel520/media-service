package tg

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) V1AdminTgsPost(ctx context.Context, req *api.TgPost) (api.V1AdminTgsPostRes, error) {
	// Создание нового тг
	return &api.Tg{}, nil
}
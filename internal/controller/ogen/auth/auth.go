package auth

import (
	"context"

	api "github.com/himmel520/uoffer/mediaAd/api/oas"
)

func (h *Handler) HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error) {
	// Реализация аутентификации администратора
	return ctx, nil
}

package auth

import (
	"context"
	"errors"

	api "github.com/himmel520/media-service/api/oas"
)

func (h *Handler) HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error) {
	role, err := h.uc.GetUserRoleFromToken(t.GetToken())
	if err != nil {
		return ctx, err
	}

	if !h.uc.IsUserAdmin(role){
		return nil, errors.New("not enough permissions")
	}
	
	return ctx, nil
}

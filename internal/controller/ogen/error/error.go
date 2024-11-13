package errctrl

import (
	"context"
	"errors"
	"net/http"

	api "github.com/himmel520/media-service/api/oas"
	"github.com/ogen-go/ogen/ogenerrors"
)

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	statusCode := http.StatusInternalServerError
	message := err.Error()

	var ogenSecurityError *ogenerrors.SecurityError
	if errors.As(err, &ogenSecurityError) {
		statusCode = ogenSecurityError.Code()
		message = ogenSecurityError.Error()
	}

	return &api.ErrorStatusCode{
		StatusCode: statusCode,
		Response: api.Error{
			Message: message,
			Details: nil,
		},
	}
}

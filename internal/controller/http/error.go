package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpSignalError interface {
	error
	Status() int
}

type errorResponse struct {
	Message string `json:"message"`
}

type HttpError struct {
	message string
	status  int
}

func NewHttpError(message string, status int) *HttpError {
	return &HttpError{
		message: message,
		status:  status,
	}
}

func (e *HttpError) Error() string {
	return e.message
}

func (e *HttpError) Status() int {
	return e.status
}

var (
	ErrInvalidID         = NewHttpError("invalid id", http.StatusBadRequest)
	ErrEmptyAuthHeader   = NewHttpError("authorization header is missing", http.StatusUnauthorized)
	ErrInvalidAuthHeader = NewHttpError("authorization header is invalid", http.StatusUnauthorized)
	ErrForbidden         = NewHttpError("you don't have access to this resource", http.StatusForbidden)
)

func checkHttpErr(h *Handler, c *gin.Context, err error, signalErrors []HttpSignalError) {
	for _, sigerr := range signalErrors {
		if errors.Is(err, sigerr) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
		}

	}
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})

}

func wrapToHttpErr(errs []error, statuses []int) []HttpSignalError {
	if len(errs) != len(statuses) {
		return nil
	}

	var wrappedErrors []HttpSignalError
	for i, err := range errs {
		wrappedError := NewHttpError(err.Error(), statuses[i])
		wrappedErrors = append(wrappedErrors, wrappedError)
	}

	return wrappedErrors
}

package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
)

type HttpSignalError interface {
	error
	Status() int
	UnWrap() error
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

type HttpError struct {
	error
	status int
}

func NewHttpError(error error, status int) *HttpError {
	return &HttpError{
		error:  error,
		status: status,
	}
}

func (e *HttpError) Error() string {
	return e.error.Error()
}

func (e *HttpError) Status() int {
	return e.status
}

func (e *HttpError) UnWrap() error {
	return e.error
}

var (
	ErrInvalidID         = NewHttpError(errors.New("invalid id"), http.StatusBadRequest)
	ErrEmptyAuthHeader   = NewHttpError(errors.New("authorization header is missing"), http.StatusUnauthorized)
	ErrInvalidAuthHeader = NewHttpError(errors.New("authorization header is invalid"), http.StatusUnauthorized)
	ErrForbidden         = NewHttpError(errors.New("you don't have access to this resource"), http.StatusForbidden)

	ErrLogoNotFound   = NewHttpError(repoerr.ErrLogoNotFound, http.StatusNotFound)
	ErrLogoExist      = NewHttpError(repoerr.ErrLogoExist, http.StatusBadRequest)
	ErrLogoDependency = NewHttpError(repoerr.ErrLogoDependency, http.StatusConflict)
)

func checkHttpErr(h *Handler, c *gin.Context, err error, signalErrors []HttpSignalError) {
	for _, sigerr := range signalErrors {

		if errors.Is(err, sigerr.UnWrap()) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
			return
		}

	}
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})

}

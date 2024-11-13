package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/media-service/internal/controller"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
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

// Middleware
var (
	ErrInvalidID         = NewHttpError(controller.ErrInvalidID, http.StatusBadRequest)
	ErrEmptyAuthHeader   = NewHttpError(controller.ErrEmptyAuthHeader, http.StatusUnauthorized)
	ErrInvalidAuthHeader = NewHttpError(controller.ErrInvalidAuthHeader, http.StatusUnauthorized)
	ErrForbidden         = NewHttpError(controller.ErrForbidden, http.StatusForbidden)
)

// Logo
var (
	ErrLogoNotFound   = NewHttpError(repoerr.ErrLogoNotFound, http.StatusNotFound)
	ErrLogoExist      = NewHttpError(repoerr.ErrLogoExist, http.StatusBadRequest)
	ErrLogoDependency = NewHttpError(repoerr.ErrLogoDependency, http.StatusConflict)
)

// Adv
var (
	ErrAdvNotFound           = NewHttpError(repoerr.ErrAdvNotFound, http.StatusNotFound)
	ErrAdvDependencyNotExist = NewHttpError(repoerr.ErrAdvDependencyNotExist, http.StatusConflict)
)

// Color
var (
	ErrColorNotFound        = NewHttpError(repoerr.ErrColorNotFound, http.StatusNotFound)
	ErrColorHexExist        = NewHttpError(repoerr.ErrColorHexExist, http.StatusBadRequest)
	ErrColorDependencyExist = NewHttpError(repoerr.ErrColorDependencyExist, http.StatusConflict)
)

// TG
var (
	ErrTGNotFound        = NewHttpError(repoerr.ErrTGNotFound, http.StatusNotFound)
	ErrTGExist           = NewHttpError(repoerr.ErrTGExist, http.StatusBadRequest)
	ErrTGDependencyExist = NewHttpError(repoerr.ErrTGDependencyExist, http.StatusConflict)
)

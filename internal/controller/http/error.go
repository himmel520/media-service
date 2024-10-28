package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/controller"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
)

type errorResponse struct {
	Message string `json:"message"`
}

func checkRepoErr(h *Handler, c *gin.Context, err error) {
	switch {
	case errors.Is(err, repoerr.ErrLogoExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return

	case errors.Is(err, repoerr.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return

	case errors.Is(err, repoerr.ErrLogoDependency):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return

	}
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})

}

func checkHttpErr(h *Handler, c *gin.Context, err error, signalErrors []controller.SignalError) {
	for _, sigerr := range signalErrors {
		if errors.Is(err, sigerr) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
		}

	}
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})

}

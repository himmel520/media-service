package httpctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/controller"
)

func checkErr(h *Handler, c *gin.Context, err error, signalErrors []controller.SignalError) {
	for _, sigerr := range signalErrors {
		if errors.Is(err, sigerr) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
			return
		}

	}
	h.log.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})

	return

}

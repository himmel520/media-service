package httpctrl

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) validateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"invalid id"})
			return
		}

		if id <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"invalid id"})
			return
		}
	}
}


func (h *Handler) jwtAdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"Authorization header is missing"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" || token == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"Authorization header is invalid"})
			return
		}

		userRole, err := h.uc.Auth.GetUserRoleFromToken(token)
		if err != nil {
			h.log.Info(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{err.Error()})
			return
		}

		if !h.uc.Auth.IsUserAdmin(userRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse{"You don't have access to this resource"})
			return
		}
	}
}

func (h *Handler) deleteCategoriesCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Исключаем метод Get, тк не вносит никаких изменений
		if c.Request.Method == http.MethodGet {
			return
		}

		// выполнение любого CRUD
		c.Next()

		// если ошибка - выход
		if c.IsAborted() {
			return
		}

		go func() {
			// удаление кэша в фоне
			if err := h.uc.Adv.DeleteCache(context.Background()); err != nil {
				h.log.Error(err)
			}
		}()
	}
}

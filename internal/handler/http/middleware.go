package httphandler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/models"
)

func (h *Handler) validateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := strconv.Atoi(c.Param("id")); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"invalid id"})
			return
		}
	}
}

// anonym - доступ с ограничением user интерфейса
// user - полный доступ к user интерфейсу
// admin - доступ ко всему сайту
func (h *Handler) jwtAuthAccess(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			if requiredRole == "admin" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"Authorization header is missing"})
				return
			}
			c.Set("role", models.RoleAnonym)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" || token == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"Authorization header is invalid"})
			return
		}

		userRole, err := h.authSrv.GetUserRoleFromToken(token)
		if err != nil {
			h.log.Info(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{err.Error()})
			return
		}

		if !h.authSrv.IsUserAuthorized(requiredRole, userRole) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"You don't have access to this resource"})
			return
		}

		c.Set("role", userRole)
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
			if err := h.advSrv.DeleteCache(context.Background()); err != nil {
				h.log.Error(err)
			}
		}()
	}
}

package httpctrl

// import (
// 	"context"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func (h *Handler) newCors() gin.HandlerFunc {
// 	cfg := cors.DefaultConfig()
// 	cfg.AllowOrigins = []string{
// 		"http://localhost",
// 		"http://localhost:5173",
// 	}
// 	cfg.AllowCredentials = true
// 	return cors.New(cfg)
// }

// func (h *Handler) validateID() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(ErrInvalidID.status, errorResponse{ErrInvalidID.Error()})
// 			return
// 		}

// 		if id <= 0 {
// 			c.AbortWithStatusJSON(ErrInvalidID.status, errorResponse{ErrInvalidID.Error()})
// 			return
// 		}
// 	}
// }

// func (h *Handler) jwtAdminAccess() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(ErrEmptyAuthHeader.status, errorResponse{ErrEmptyAuthHeader.Error()})
// 			return
// 		}

// 		token := strings.TrimPrefix(authHeader, "Bearer ")
// 		if token == "" || token == authHeader {
// 			c.AbortWithStatusJSON(ErrInvalidAuthHeader.status, errorResponse{ErrInvalidAuthHeader.Error()})
// 			return
// 		}

// 		userRole, err := h.uc.Auth.GetUserRoleFromToken(token)
// 		if err != nil {
// 			h.log.Info(err.Error())
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{err.Error()})
// 			return
// 		}

// 		if !h.uc.Auth.IsUserAdmin(userRole) {
// 			c.AbortWithStatusJSON(ErrForbidden.status, errorResponse{ErrForbidden.Error()})
// 			return
// 		}
// 	}
// }

// func (h *Handler) deleteCategoriesCache() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Исключаем метод Get, тк не вносит никаких изменений
// 		if c.Request.Method == http.MethodGet {
// 			return
// 		}

// 		// выполнение любого CRUD
// 		c.Next()

// 		// если ошибка - выход
// 		if c.IsAborted() {
// 			return
// 		}

// 		if err := h.uc.Adv.DeleteCache(context.Background()); err != nil {
// 			h.log.Error(err)
// 		}
// 	}
// }

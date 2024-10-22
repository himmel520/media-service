package httpctrl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
)

// @Summary Добавить новый цвет
// @Description Создает новый цвет
// @Tags colors
// @Accept json
// @Produce json
// @Param color body entity.Color true "Цвет"
// @Success 201 {object} entity.ColorResp "Успешное создание цвета"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/colors [post]
func (h *Handler) addColor(c *gin.Context) {
	var color *entity.Color
	if err := c.BindJSON(&color); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newColor, err := h.uc.Color.Add(c.Request.Context(), color)
	switch {
	case errors.Is(err, repoerr.ErrColorHexExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return

	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return

	}

	c.JSON(http.StatusCreated, newColor)
}

// @Summary Обновить цвет
// @Description Обновляет существующий цвет
// @Tags colors
// @Accept json
// @Produce json
// @Param id path int true "id цвета"
// @Param color body entity.ColorUpdate true "Обновленный цвет"
// @Success 200 {object} entity.ColorResp "Успешное обновление цвета"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Цвет не найден"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/colors/{id} [put]
func (h *Handler) updateColor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var color *entity.ColorUpdate
	if err := c.BindJSON(&color); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if color.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"color has no changes"})
		return
	}

	newColor, err := h.uc.Color.Update(c.Request.Context(), id, color)
	switch {
	case errors.Is(err, repoerr.ErrColorHexExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	case errors.Is(err, repoerr.ErrColorNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, newColor)
}

// @Summary Удалить цвет
// @Description Удаляет цвет по его идентификатору
// @Tags colors
// @Produce json
// @Param id path int true "id цвета"
// @Success 204 "Успешное удаление"
// @Failure 404 {object} errorResponse "Цвет не найден"
// @Failure 409 {object} errorResponse "Конфликт данных"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/colors/{id} [delete]
func (h *Handler) deleteColor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.uc.Color.Delete(c.Request.Context(), id)
	switch {
	case errors.Is(err, repoerr.ErrColorDependencyExist):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case errors.Is(err, repoerr.ErrColorNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Получить список цветов
// @Description Возвращает список всех цветов с учетом пагинации
// @Tags colors
// @Produce json
// @Param query query PaginationQuery true "Параметры пагинации"
// @Success 200 {object} entity.ColorsResp "Список цветов"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Цвета не найдены"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /colors [get]
func (h *Handler) getColors(c *gin.Context) {
	var query *PaginationQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{fmt.Sprintf("неккоректный query: %v", err)})
		return
	}

	colors, err := h.uc.Color.GetAllWithPagination(c.Request.Context(), query.Limit, query.Offset)
	switch {
	case errors.Is(err, repoerr.ErrColorNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, colors)
}

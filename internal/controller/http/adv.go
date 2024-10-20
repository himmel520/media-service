package httpctrl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
)

// @Summary Добавить новое объявление
// @Description Создает новое объявление на основе переданных данных
// @Tags adv
// @Accept json
// @Produce json
// @Param adv body entity.Adv true "Объявление"
// @Success 201 {object} entity.AdvResponse "Успешное создание объявления"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 409 {object} errorResponse "Конфликт данных"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/ads [post]
func (h *Handler) addAdv(c *gin.Context) {
	var adv *entity.Adv
	if err := c.BindJSON(&adv); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	advResp, err := h.srv.Adv.Add(c.Request.Context(), adv)
	switch {
	case errors.Is(err, repoerr.ErrAdvDependencyNotExist):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case err != nil:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, advResp)
}

// @Summary Удалить объявление
// @Description Удаляет объявление по его id
// @Tags adv
// @Produce json
// @Param id path int true "id объявления"
// @Success 204 "Успешное удаление"
// @Failure 404 {object} errorResponse "Объявление не найдено"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/ads/{id} [delete]
func (h *Handler) deleteAdv(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.srv.Adv.Delete(c.Request.Context(), id)
	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Обновить объявление
// @Description Обновляет существующее объявление
// @Tags adv
// @Accept json
// @Produce json
// @Param id path int true "id объявления"
// @Param adv body entity.AdvUpdate true "Обновленное объявление"
// @Success 200 {object} entity.AdvResponse "Успешное обновление объявления"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Объявление не найдено"
// @Failure 409 {object} errorResponse "Конфликт данных"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/ads/{id} [put]
func (h *Handler) updateAdv(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var adv *entity.AdvUpdate
	if err := c.BindJSON(&adv); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if adv.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"no changes"})
		return
	}

	advResp, err := h.srv.Adv.Update(c.Request.Context(), id, adv)
	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case errors.Is(err, repoerr.ErrAdvDependencyNotExist):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case err != nil:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, advResp)
}

// @Summary Получить объявления с фильтрацией
// @Description Возвращает список объявлений с возможностью фильтрации
// @Tags adv
// @Produce json
// @Param query query AdvPostQuery true "Параметры фильтрации"
// @Success 200 {array} entity.AdvResponse "Список объявлений"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Объявления не найдены"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /ads [get]
func (h *Handler) getAdvsWithFilter(c *gin.Context) {
	var query AdvPostQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if len(query.Priority) == 0 {
		query.SetDefaultPriority()
	}

	advs, err := h.srv.Adv.GetAllWithFilter(c.Request.Context(), query.Limit, query.Offset, query.Post, query.Priority)
	switch {
	case errors.Is(err, repoerr.ErrAdvNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, advs)
}

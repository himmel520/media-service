package httphandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/repository"
	"github.com/himmel520/uoffer/mediaAd/models"
)

// @Summary Добавить группу Telegram
// @Description Создает новую группу Telegram
// @Tags tg
// @Accept json
// @Produce json
// @Param tg body models.TG true "Группа Telegram"
// @Success 201 {object} models.TGResp "Успешное создание группы"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs [post]
func (h *Handler) addTG(c *gin.Context) {
	var tg *models.TG
	if err := c.BindJSON(&tg); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newTG, err := h.srv.AddTG(c.Request.Context(), tg)
	if err != nil {
		if errors.Is(err, repository.ErrTGExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}

		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTG)
}

// @Summary Получить список групп Telegram
// @Description Возвращает список всех групп Telegram с учетом пагинации
// @Tags tg
// @Produce json
// @Param query query PaginationQuery true "Параметры пагинации"
// @Success 200 {object} models.TGsResp "Список групп Telegram"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Группы Telegram не найдены"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs [get]
func (h *Handler) getTGs(c *gin.Context) {
	var query *PaginationQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{fmt.Sprintf("неккоректный query: %v", err)})
		return
	}

	tgs, err := h.srv.GetTGs(c.Request.Context(), query.Limit, query.Offset)
	switch {
	case errors.Is(err, repository.ErrTGNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, tgs)
}

// @Summary Обновить группу Telegram
// @Description Обновляет существующую группу Telegram
// @Tags tg
// @Accept json
// @Produce json
// @Param id path int true "id группы Telegram"
// @Param tg body models.TGUpdate true "Обновленная группа Telegram"
// @Success 200 {object} models.TGResp "Успешное обновление группы"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Группа Telegram не найдена"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs/{id} [put]
func (h *Handler) updateTG(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tg *models.TGUpdate
	if err := c.BindJSON(&tg); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if tg.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"TG has no changes"})
		return
	}

	newTG, err := h.srv.UpdateTG(c.Request.Context(), id, tg)
	switch {
	case errors.Is(err, repository.ErrTGExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	case errors.Is(err, repository.ErrTGNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTG)
}

// @Summary Удалить группу Telegram
// @Description Удаляет группу Telegram по ее id
// @Tags tg
// @Produce json
// @Param id path int true "idгруппы Telegram"
// @Success 204 "Успешное удаление"
// @Failure 404 {object} errorResponse "Группа Telegram не найдена"
// @Failure 409 {object} errorResponse "Конфликт данных"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs/{id} [delete]
func (h *Handler) deleteTG(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.srv.DeleteTG(c.Request.Context(), id)
	switch {
	case errors.Is(err, repository.ErrTGDependencyExist):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case errors.Is(err, repository.ErrTGNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
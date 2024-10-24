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

// @Summary Добавить группу Telegram
// @Description Создает новую группу Telegram
// @Tags tg
// @Accept json
// @Produce json
// @Param tg body entity.TG true "Группа Telegram"
// @Success 201 {object} entity.TGResp "Успешное создание группы"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs [post]
func (h *Handler) addTG(c *gin.Context) {
	var tg *entity.TG
	if err := c.BindJSON(&tg); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newTG, err := h.uc.TG.Add(c.Request.Context(), tg)
	if err != nil {
		if errors.Is(err, repoerr.ErrTGExist) {
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
// @Success 200 {object} entity.TGsResp "Список групп Telegram"
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

	tgs, err := h.uc.TG.GetAllWithPagination(c.Request.Context(), query.Limit, query.Offset)
	switch {
	case errors.Is(err, repoerr.ErrTGNotFound):
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
// @Param tg body entity.TGUpdate true "Обновленная группа Telegram"
// @Success 200 {object} entity.TGResp "Успешное обновление группы"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Группа Telegram не найдена"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/tgs/{id} [put]
func (h *Handler) updateTG(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tg *entity.TGUpdate
	if err := c.BindJSON(&tg); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if tg.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"TG has no changes"})
		return
	}

	newTG, err := h.uc.TG.Update(c.Request.Context(), id, tg)
	switch {
	case errors.Is(err, repoerr.ErrTGExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	case errors.Is(err, repoerr.ErrTGNotFound):
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

	err := h.uc.TG.Delete(c.Request.Context(), id)
	switch {
	case errors.Is(err, repoerr.ErrTGDependencyExist):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case errors.Is(err, repoerr.ErrTGNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

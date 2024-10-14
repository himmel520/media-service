package httphandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/uoffer/mediaAd/internal/repository"
	"github.com/himmel520/uoffer/mediaAd/internal/models"
)

// @Summary Добавить новый логотип
// @Description Создает новый логотип
// @Tags logos
// @Accept json
// @Produce json
// @Param logo body models.Logo true "Логотип"
// @Success 201 {object} models.LogoResp "Успешное создание логотипа"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/logos [post]
func (h *Handler) addLogo(c *gin.Context) {
	var logo *models.Logo
	if err := c.BindJSON(&logo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newLogo, err := h.logoSrv.Add(c.Request.Context(), logo)
	if err != nil {
		if errors.Is(err, repository.ErrLogoExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}

		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLogo)
}

// @Summary Обновить логотип
// @Description Обновляет существующий логотип
// @Tags logos
// @Accept json
// @Produce json
// @Param id path int true "id логотипа"
// @Param logo body models.LogoUpdate true "Обновленный логотип"
// @Success 200 {object} models.LogoResp "Успешное обновление логотипа"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Логотип не найден"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/logos/{id} [put]
func (h *Handler) updateLogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var logo *models.LogoUpdate
	if err := c.BindJSON(&logo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if logo.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"logo has no changes"})
		return
	}

	newLogo, err := h.logoSrv.Update(c.Request.Context(), id, logo)
	switch {
	case errors.Is(err, repository.ErrLogoExist):
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	case errors.Is(err, repository.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, newLogo)
}

// @Summary Удалить логотип
// @Description Удаляет логотип по его id
// @Tags logos
// @Produce json
// @Param id path int true "id логотипа"
// @Success 204 "Успешное удаление"
// @Failure 404 {object} errorResponse "Логотип не найден"
// @Failure 409 {object} errorResponse "Конфликт данных"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/logos/{id} [delete]
func (h *Handler) deleteLogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.logoSrv.Delete(c.Request.Context(), id)
	switch {
	case errors.Is(err, repository.ErrLogoDependency):
		c.AbortWithStatusJSON(http.StatusConflict, errorResponse{err.Error()})
		return
	case errors.Is(err, repository.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Получить список логотипов
// @Description Возвращает список всех логотипов с учетом пагинации.
// @Tags logos
// @Produce json
// @Param query query PaginationQuery true "Параметры пагинации"
// @Success 200 {object} models.LogosResp "Список логотипов"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Логотипы не найдены"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /logos [get]
func (h *Handler) getPaginatedLogos(c *gin.Context) {
	var query *PaginationQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{fmt.Sprintf("неккоректный query: %v", err)})
		return
	}

	logos, err := h.logoSrv.GetAllWithPagination(c.Request.Context(), query.Limit, query.Offset)
	switch {
	case errors.Is(err, repository.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, logos)
}

func (h *Handler) getLogos(c *gin.Context) {
	logos, err := h.logoSrv.GetAll(c.Request.Context())
	switch {
	case errors.Is(err, repository.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, logos)
}

// @Summary Получить логотип
// @Description Возвращает логотип по его id
// @Tags logos
// @Produce json
// @Param id path int true "id логотипа"
// @Success 200 {object} models.LogoResp "Логотип"
// @Failure 404 {object} errorResponse "Логотип не найден"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /logos/{id} [get]
func (h *Handler) getLogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	logo, err := h.logoSrv.GetByID(c.Request.Context(), id)
	switch {
	case errors.Is(err, repository.ErrLogoNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	case err != nil:
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, logo)
}

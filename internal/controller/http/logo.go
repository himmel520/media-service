package httpctrl

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	con "github.com/himmel520/uoffer/mediaAd/internal/controller"
	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/repoerr"
)

// @Summary Добавить новый логотип
// @Description Создает новый логотип
// @Tags logos
// @Accept json
// @Produce json
// @Param logo body entity.Logo true "Логотип"
// @Success 201 {object} entity.LogoResp "Успешное создание логотипа"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/logos [post]
func (h *Handler) addLogo(c *gin.Context) {
	var logo *entity.Logo
	if err := c.BindJSON(&logo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	newLogo, err := h.uc.Logo.Add(c.Request.Context(), logo)
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoExist})
	}

	c.JSON(http.StatusCreated, newLogo)
}

// @Summary Обновить логотип
// @Description Обновляет существующий логотип
// @Tags logos
// @Accept json
// @Produce json
// @Param id path int true "id логотипа"
// @Param logo body entity.LogoUpdate true "Обновленный логотип"
// @Success 200 {object} entity.LogoResp "Успешное обновление логотипа"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Логотип не найден"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /admin/logos/{id} [put]
func (h *Handler) updateLogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var logo *entity.LogoUpdate
	if err := c.BindJSON(&logo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	if logo.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{"logo has no changes"})
		return
	}

	newLogo, err := h.uc.Logo.Update(c.Request.Context(), id, logo)
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoExist, repoerr.ErrLogoNotFound})
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

	err := h.uc.Logo.Delete(c.Request.Context(), id)
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoDependency, repoerr.ErrLogoNotFound})
	}

	c.Status(http.StatusNoContent)
}

// @Summary Получить список логотипов
// @Description Возвращает список всех логотипов с учетом пагинации.
// @Tags logos
// @Produce json
// @Param query query PaginationQuery true "Параметры пагинации"
// @Success 200 {object} entity.LogosResp "Список логотипов"
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

	logos, err := h.uc.Logo.GetAllWithPagination(c.Request.Context(), query.Limit, query.Offset)
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoNotFound})
	}

	c.JSON(http.StatusOK, logos)
}

// @Summary Получить список логотипов
// @Description Возвращает список всех логотипов
// @Tags logos
// @Produce json
// @Success 200 {object} entity.LogosResp "Список логотипов"
// @Failure 400 {object} errorResponse "Неверные данные"
// @Failure 404 {object} errorResponse "Логотипы не найдены"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /logos [get]
func (h *Handler) getLogos(c *gin.Context) {
	logos, err := h.uc.Logo.GetAll(c.Request.Context())
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoNotFound})
	}

	c.JSON(http.StatusOK, logos)
}

// @Summary Получить логотип
// @Description Возвращает логотип по его id
// @Tags logos
// @Produce json
// @Param id path int true "id логотипа"
// @Success 200 {object} entity.LogoResp "Логотип"
// @Failure 404 {object} errorResponse "Логотип не найден"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /logos/{id} [get]
func (h *Handler) getLogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	logo, err := h.uc.Logo.GetByID(c.Request.Context(), id)
	if err != nil {
		checkErr(h, c, err, []con.SignalError{repoerr.ErrLogoNotFound})
	}

	c.JSON(http.StatusOK, logo)
}

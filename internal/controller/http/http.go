package httpctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/himmel520/media-service/internal/usecase"
	// _ "github.com/himmel520/uoffer/mediaAd/docs"

	"github.com/sirupsen/logrus"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	uc  *usecase.Usecase
	log *logrus.Logger
}

func NewHandler(uc *usecase.Usecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(h.newCors())

	api := r.Group("/api/v1")
	{
		// api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		logo := api.Group("/logos")
		{
			logo.GET("/", h.getLogos)

			logo.Use(h.validateID())
			logo.GET("/:id", h.getLogo)
		}

		adv := api.Group("/ads")
		{
			adv.GET("/", h.getAdvsWithFilter)
		}

		// admin
		admin := api.Group("/admin", h.jwtAdminAccess())
		{
			logo := admin.Group("/logos")
			{
				logo.GET("/", h.getPaginatedLogos)
				logo.POST("/", h.addLogo)

				logo.Use(h.validateID())
				logo.PUT("/:id", h.updateLogo)
				logo.DELETE("/:id", h.deleteLogo)
			}

			colors := admin.Group("/colors", h.deleteCategoriesCache())
			{
				colors.POST("/", h.addColor)
				colors.GET("/", h.getColors)

				colors.Use(h.validateID())
				colors.DELETE("/:id", h.deleteColor)
				colors.PUT("/:id", h.updateColor)
			}

			tg := admin.Group("/tgs", h.deleteCategoriesCache())
			{
				tg.POST("/", h.addTG)
				tg.GET("/", h.getTGs)

				tg.Use(h.validateID())
				tg.DELETE("/:id", h.deleteTG)
				tg.PUT("/:id", h.updateTG)
			}

			adv := admin.Group("/ads", h.deleteCategoriesCache())
			{
				adv.POST("/", h.addAdv)

				adv.Use(h.validateID())
				adv.PUT("/:id", h.updateAdv)
				adv.DELETE("/:id", h.deleteAdv)
			}
		}
	}

	return r
}

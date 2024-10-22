package httpctrl

import (
	"github.com/gin-gonic/gin"
	_ "github.com/himmel520/uoffer/mediaAd/docs"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase"

	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	uc  *usecase.Usecase
	log *logrus.Logger
}

func New(uc *usecase.Usecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		logo := api.Group("/logos")
		{
			logo.GET("/", h.getLogos) // get logos

			logo.Use(h.validateID())
			logo.GET("/:id", h.getLogo) // get logo by id
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
				logo.GET("/", h.getPaginatedLogos) // get logos
				logo.POST("/", h.addLogo)          // Add a new logo

				logo.Use(h.validateID())
				logo.PUT("/:id", h.updateLogo)    // Update a logo
				logo.DELETE("/:id", h.deleteLogo) // Delete a logo
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

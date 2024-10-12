package httphandler

import (
	"context"
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	_ "github.com/himmel520/uoffer/mediaAd/docs"
	"github.com/himmel520/uoffer/mediaAd/internal/config"
	"github.com/himmel520/uoffer/mediaAd/models"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	Service interface {
		Auth
		Cache
		Logo

		Colors
		Tg
		Adv
	}

	Auth interface {
		GetUserRoleFromToken(jwtToken string, publicKey *rsa.PublicKey) (string, error)
		IsUserAuthorized(requiredRole, userRole string) bool
	}

	Logo interface {
		AddLogo(ctx context.Context, logo *models.Logo) (*models.LogoResp, error)
		UpdateLogo(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error)
		DeleteLogo(ctx context.Context, id int) error
		GetLogo(ctx context.Context, logoID int) (*models.LogoResp, error)
		GetLogos(ctx context.Context, limit, offset int) (*models.LogosResp, error)
	}

	Colors interface {
		AddColor(ctx context.Context, color *models.Color) (*models.ColorResp, error)
		UpdateColor(ctx context.Context, id int, color *models.ColorUpdate) (*models.ColorResp, error)
		DeleteColor(ctx context.Context, id int) error
		GetColors(ctx context.Context, limit, offset int) (*models.ColorsResp, error)
	}

	Tg interface {
		AddTG(ctx context.Context, tg *models.TG) (*models.TGResp, error)
		UpdateTG(ctx context.Context, id int, TG *models.TGUpdate) (*models.TGResp, error)
		DeleteTG(ctx context.Context, id int) error
		GetTGs(ctx context.Context, limit, offset int) (*models.TGsResp, error)
	}

	Adv interface {
		AddAdv(ctx context.Context, adv *models.Adv) (*models.AdvResponse, error)
		DeleteAdv(ctx context.Context, id int) error
		UpdateAdv(ctx context.Context, id int, adv *models.AdvUpdate) (*models.AdvResponse, error)
		GetAdvsWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error)
	}

	Cache interface {
		DeleteAdvsCache(ctx context.Context) error
	}
)

type Handler struct {
	srv Service
	log *logrus.Logger
	cfg *config.JWT
}

func New(srv Service, cfg *config.JWT, log *logrus.Logger) *Handler {
	return &Handler{srv: srv, log: log, cfg: cfg}
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
		admin := api.Group("/admin", h.jwtAuthAccess(models.RoleAdmin))
		{
			logo := admin.Group("/logos")
			{
				logo.POST("/", h.addLogo) // Add a new logo

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

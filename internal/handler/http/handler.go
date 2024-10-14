package httphandler

import (
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/himmel520/uoffer/mediaAd/docs"
	"github.com/himmel520/uoffer/mediaAd/internal/models"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate mockery --all

type (
	AdvSrv interface {
		Add(ctx context.Context, adv *models.Adv) (*models.AdvResponse, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, adv *models.AdvUpdate) (*models.AdvResponse, error)
		GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error)
		DeleteCache(ctx context.Context) error
	}

	AuthSrv interface {
		GetUserRoleFromToken(jwtToken string) (string, error)
		IsUserAuthorized(requiredRole, userRole string) bool
	}

	ColorSrv interface {
		Add(ctx context.Context, color *models.Color) (*models.ColorResp, error)
		Update(ctx context.Context, id int, color *models.ColorUpdate) (*models.ColorResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) (*models.ColorsResp, error)
	}

	LogoSrv interface {
		Add(ctx context.Context, logo *models.Logo) (*models.LogoResp, error)
		Update(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error)
		Delete(ctx context.Context, id int) error
		GetByID(ctx context.Context, id int) (*models.LogoResp, error)
		GetAll(ctx context.Context) ([]*models.Logo, error)
		GetAllWithPagination(ctx context.Context, limit, offset int) (*models.LogosResp, error)
	}

	TGSrv interface {
		Add(ctx context.Context, tg *models.TG) (*models.TGResp, error)
		Update(ctx context.Context, id int, TG *models.TGUpdate) (*models.TGResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) (*models.TGsResp, error)
	}
)

type Handler struct {
	advSrv   AdvSrv
	authSrv  AuthSrv
	colorSrv ColorSrv
	logoSrv  LogoSrv
	tgSrv    TGSrv
	log      *logrus.Logger
}

func New(advSrv AdvSrv, authSrv AuthSrv, colorSrv ColorSrv, logoSrv LogoSrv, tgSrv TGSrv, log *logrus.Logger) *Handler {
	return &Handler{
		advSrv:   advSrv,
		authSrv:  authSrv,
		colorSrv: colorSrv,
		logoSrv:  logoSrv,
		tgSrv:    tgSrv,
		log:      log,
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
		admin := api.Group("/admin", h.jwtAuthAccess(models.RoleAdmin))
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

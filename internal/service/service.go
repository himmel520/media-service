package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/models"
	"github.com/sirupsen/logrus"
)

type (
	Repo interface {
		Logo
		Color
		TG
		Adv
	}

	Logo interface {
		AddLogo(ctx context.Context, logo *models.Logo) (*models.LogoResp, error)
		UpdateLogo(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error)
		DeleteLogo(ctx context.Context, id int) error
		GetLogo(ctx context.Context, logoID int) (*models.LogoResp, error)
		GetLogos(ctx context.Context, limit, offset int) (map[int]*models.Logo, error)
		GetLogoCount(ctx context.Context) (int, error)
	}

	Color interface {
		AddColor(ctx context.Context, color *models.Color) (*models.ColorResp, error)
		UpdateColor(ctx context.Context, id int, Color *models.ColorUpdate) (*models.ColorResp, error)
		DeleteColor(ctx context.Context, id int) error
		GetColors(ctx context.Context, limit, offset int) ([]*models.ColorResp, error)
		GetColorCount(ctx context.Context) (int, error)
	}

	TG interface {
		AddTG(ctx context.Context, tg *models.TG) (*models.TGResp, error)
		UpdateTG(ctx context.Context, id int, tg *models.TGUpdate) (*models.TGResp, error)
		DeleteTG(ctx context.Context, id int) error
		GetTGs(ctx context.Context, limit, offset int) ([]*models.TGResp, error)
		GetTGCount(ctx context.Context) (int, error)
	}

	Adv interface {
		AddAdv(ctx context.Context, adv *models.Adv) (int, error)
		GetAdvByID(ctx context.Context, id int) (*models.AdvResponse, error)
		DeleteAdv(ctx context.Context, id int) error
		UpdateAdv(ctx context.Context, id int, adv *models.AdvUpdate) error
		GetAdvsWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error)
	}

	Cache interface {
		SetAdv(ctx context.Context, key string, advs []*models.AdvResponse) error
		GetAdv(ctx context.Context, key string) ([]*models.AdvResponse, error)
		DeleteAdvsCache(ctx context.Context) error
	}

	Service struct {
		repo  Repo
		cache Cache
		log   *logrus.Logger
	}
)

func New(repo Repo, cache Cache, log *logrus.Logger) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		log:   log,
	}
}

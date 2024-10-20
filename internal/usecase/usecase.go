package usecase

import (
	"context"
	"crypto/rsa"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"

	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --all

type (
	Usecase struct {
		Adv   AdvSrv
		Auth  AuthSrv
		Color ColorSrv
		Logo  LogoSrv
		TG    TGSrv
	}

	AdvSrv interface {
		Add(ctx context.Context, adv *entity.Adv) (*entity.AdvResponse, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResponse, error)
		GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error)
		DeleteCache(ctx context.Context) error
	}

	AuthSrv interface {
		GetUserRoleFromToken(jwtToken string) (string, error)
		IsUserAuthorized(requiredRole, userRole string) bool
	}

	ColorSrv interface {
		Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error)
		Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.ColorResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.ColorsResp, error)
	}

	LogoSrv interface {
		Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error)
		Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error)
		Delete(ctx context.Context, id int) error
		GetByID(ctx context.Context, id int) (*entity.LogoResp, error)
		GetAll(ctx context.Context) ([]*entity.LogoResp, error)
		GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.LogosResp, error)
	}

	TGSrv interface {
		Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error)
		Update(ctx context.Context, id int, TG *entity.TGUpdate) (*entity.TGResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.TGsResp, error)
	}
)

func New(repo *repository.Repository, cache *cache.Cache, publicKey rsa.PublicKey, log *logrus.Logger) *Usecase {
	return &Usecase{
		Adv:   NewAdvUsecase(repo.AdvRepo, cache.Adv, log),
		Auth:  NewAuthUsecase(publicKey),
		Color: NewColorUsecase(repo.ColorRepo, log),
		Logo:  NewLogoUsecase(repo.LogoRepo, log),
		TG:    NewTGUsecase(repo.TGRepo, log),
	}
}

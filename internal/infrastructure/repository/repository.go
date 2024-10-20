package repository

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Repository struct {
		AdvRepo
		ColorRepo
		LogoRepo
		TGRepo
	}

	AdvRepo interface {
		Add(ctx context.Context, adv *entity.Adv) (int, error)
		GetByID(ctx context.Context, id int) (*entity.AdvResponse, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, adv *entity.AdvUpdate) error
		GetAllWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*entity.AdvResponse, error)
	}

	ColorRepo interface {
		Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error)
		Update(ctx context.Context, id int, Color *entity.ColorUpdate) (*entity.ColorResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) ([]*entity.ColorResp, error)
		Count(ctx context.Context) (int, error)
	}

	LogoRepo interface {
		Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error)
		Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error)
		Delete(ctx context.Context, id int) error
		GetByID(ctx context.Context, logoID int) (*entity.LogoResp, error)
		GetAllWithPagination(ctx context.Context, limit, offset int) (map[int]*entity.Logo, error)
		GetAll(ctx context.Context) ([]*entity.LogoResp, error)
		Count(ctx context.Context) (int, error)
	}

	TGRepo interface {
		Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error)
		Update(ctx context.Context, id int, tg *entity.TGUpdate) (*entity.TGResp, error)
		Delete(ctx context.Context, id int) error
		GetAllWithPagination(ctx context.Context, limit, offset int) ([]*entity.TGResp, error)
		Count(ctx context.Context) (int, error)
	}
)

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		AdvRepo:   postgres.NewAdvRepo(pool),
		ColorRepo: postgres.NewColorRepo(pool),
		LogoRepo:  postgres.NewLogorRepo(pool),
		TGRepo:    postgres.NewTGRepo(pool),
	}
}

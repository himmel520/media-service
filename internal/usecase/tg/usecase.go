package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

type (
	TgUC struct {
		db   DBXT
		repo TgRepo
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	TgRepo interface {
		Create(ctx context.Context, qe repository.Querier, tg *entity.Tg) (*entity.Tg, error)
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Tg, error)
		Update(ctx context.Context, qe repository.Querier, id int, tg *entity.TgUpdate) (*entity.Tg, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}
)

func New(db DBXT, repo TgRepo) *TgUC {
	return &TgUC{
		db:   db,
		repo: repo,
	}
}

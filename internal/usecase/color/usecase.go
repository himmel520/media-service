package colorUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

type (
	ColorUC struct {
		db   DBXT
		repo ColorRepo
		cache Cache
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	ColorRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Color, error)
		Create(ctx context.Context, qe repository.Querier, color *entity.Color) (*entity.Color, error)
		Update(ctx context.Context, qe repository.Querier, id int, color *entity.ColorUpdate) (*entity.Color, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}

	Cache interface {
		Delete(ctx context.Context, prefix string) error
	}
)

func New(db DBXT, repo ColorRepo, cache Cache) *ColorUC {
	return &ColorUC{
		db:   db,
		repo: repo,
		cache: cache,
	}
}

package colorUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	ColorUC struct {
		db   DBXT
		repo ColorRepo
		log  *logrus.Logger
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
)

func New(db DBXT, repo ColorRepo, log *logrus.Logger) *ColorUC {
	return &ColorUC{
		db:   db,
		repo: repo,
		log:  log,
	}
}

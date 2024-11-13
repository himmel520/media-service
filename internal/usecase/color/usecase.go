package colorUC

import (
	"context"

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
	}
)

func New(db DBXT, repo ColorRepo, log *logrus.Logger) *ColorUC {
	return &ColorUC{
		db:   db,
		repo: repo,
		log:  log,
	}
}

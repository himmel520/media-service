package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	ImgUC struct {
		db    DBXT
		repo  ImgRepo
		cache cache.Cache
		log   *logrus.Logger
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	ImgRepo interface {
	}
)

func New(db DBXT, repo ImgRepo, cache cache.Cache, log *logrus.Logger) *ImgUC {
	return &ImgUC{db: db, repo: repo}
}

package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	AdUC struct {
		db    DBXT
		repo  AdRepo
		cache cache.Cache
		log   *logrus.Logger
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	AdRepo interface {
	}
)

func New(db DBXT, repo AdRepo, cache cache.Cache, log *logrus.Logger) *AdUC {
	return &AdUC{
		db: db, 
		repo: repo,
		cache: cache,
		log: log,
	}
}

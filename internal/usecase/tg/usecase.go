package tgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	TgUC struct {
		db   DBXT
		repo TgRepo
		log  *logrus.Logger
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	TgRepo interface {
	}
)

func New(db DBXT, repo TgRepo, log *logrus.Logger) *TgUC {
	return &TgUC{
		db:   db,
		repo: repo,
		log:  log,
	}
}

package authUC

import (
	"context"
	"crypto/rsa"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	AuthUC struct {
		publicKey rsa.PublicKey
		log       *logrus.Logger
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}
)

func New(publicKey rsa.PublicKey, log *logrus.Logger) *AuthUC {
	return &AuthUC{
		publicKey: publicKey,
		log:   log,
	}
}

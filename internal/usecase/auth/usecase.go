package authUC

import (
	"context"
	"crypto/rsa"

	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

type (
	AuthUC struct {
		publicKey rsa.PublicKey
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}
)

func New(publicKey rsa.PublicKey) *AuthUC {
	return &AuthUC{
		publicKey: publicKey,
	}
}

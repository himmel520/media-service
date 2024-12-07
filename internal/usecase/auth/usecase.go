package authUC

import (
	"crypto/rsa"
)

type (
	AuthUC struct {
		publicKey rsa.PublicKey
	}
)

func New(publicKey rsa.PublicKey) *AuthUC {
	return &AuthUC{
		publicKey: publicKey,
	}
}

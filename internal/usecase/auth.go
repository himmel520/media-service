package usecase

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/himmel520/media-service/internal/entity"
)

type AuthUsecase struct {
	publicKey rsa.PublicKey
}

func NewAuthUsecase(publicKey rsa.PublicKey) *AuthUsecase {
	return &AuthUsecase{publicKey: publicKey}
}

func (uc *AuthUsecase) GetUserRoleFromToken(jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return &uc.publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not of type jwt.MapClaims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("role claim is not a string")
	}

	return role, err
}

func (uc *AuthUsecase) IsUserAdmin(userRole string) bool {
	return userRole == entity.RoleAdmin
}

package authUC

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/himmel520/media-service/internal/entity"
)

func (uc *AuthUC) GetUserRoleFromToken(jwtToken string) (int, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return &uc.publicKey, nil
	})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("token claims are not of type jwt.MapClaims")
	}

	role, ok := claims["role"].(int)
	if !ok {
		return 0, errors.New("role claim is not a string")
	}

	return role, err
}

func (uc *AuthUC) IsUserAdmin(userRole int) bool {
	return userRole == entity.RoleAdmin
}

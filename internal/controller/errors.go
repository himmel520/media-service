package controller

import (
	"errors"
)

var (
	ErrInvalidID         = errors.New("invalid id")
	ErrEmptyAuthHeader   = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is invalid")
	ErrForbidden         = errors.New("you don't have access to this resource")
)

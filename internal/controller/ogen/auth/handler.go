package auth

import "github.com/sirupsen/logrus"

type (
	Handler struct {
		log *logrus.Logger
		uc  AuthUsecase
	}

	AuthUsecase interface{}
)

func New(uc AuthUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

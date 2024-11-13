package tg

import "github.com/sirupsen/logrus"

type (
	Handler struct {
		uc  TgUsecase
		log *logrus.Logger
	}

	TgUsecase interface{}
)

func New(uc TgUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc: uc,
		log: log,
	}
}

package color

import "github.com/sirupsen/logrus"

type (
	Handler struct {
		uc  ColorUsecase
		log *logrus.Logger
	}

	ColorUsecase interface{}
)

func New(uc ColorUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

package ad

import "github.com/sirupsen/logrus"

type (
	Handler struct {
		uc  AdUsecase
		log *logrus.Logger
	}

	AdUsecase interface{}
)

func New(uc AdUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

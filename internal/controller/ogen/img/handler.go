package img

import "github.com/sirupsen/logrus"

type (
	Handler struct {
		uc  ImgUsecase
		log *logrus.Logger
	}

	ImgUsecase interface{}
)

func New(uc ImgUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

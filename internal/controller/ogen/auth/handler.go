package auth

import "github.com/sirupsen/logrus"

type Handler struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

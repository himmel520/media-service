package tg

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/usecase"
	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		uc  TgUsecase
		log *logrus.Logger
	}

	TgUsecase interface {
		Create(ctx context.Context, tg *entity.Tg) (*entity.Tg, error)
		Get(ctx context.Context, params usecase.PageParams) (*entity.TgsResp, error)
		Update(ctx context.Context, id int, tg *entity.TgUpdate) (*entity.Tg, error)
		Delete(ctx context.Context, id int) error 
	}
)

func New(uc TgUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

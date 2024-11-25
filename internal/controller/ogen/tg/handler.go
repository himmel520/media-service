package tg

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/usecase"
)

type (
	Handler struct {
		uc TgUsecase
	}

	TgUsecase interface {
		Create(ctx context.Context, tg *entity.Tg) (*entity.Tg, error)
		Get(ctx context.Context, params usecase.PageParams) (*entity.TgsResp, error)
		Update(ctx context.Context, id int, tg *entity.TgUpdate) (*entity.Tg, error)
		Delete(ctx context.Context, id int) error
	}
)

func New(uc TgUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

package ad

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/usecase"
)

type (
	Handler struct {
		uc AdUsecase
	}

	AdUsecase interface {
		Get(ctx context.Context, params usecase.PageParams) (*entity.AdsResp, error)
		GetWithFilter(ctx context.Context, params usecase.AdvFilterParams) ([]*entity.AdvResp, error)
		Create(ctx context.Context, adv *entity.Adv) (*entity.AdvResp, error)
		Update(ctx context.Context, id int, adv *entity.AdvUpdate) (*entity.AdvResp, error)
		Delete(ctx context.Context, id int) error
	}
)

func New(uc AdUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

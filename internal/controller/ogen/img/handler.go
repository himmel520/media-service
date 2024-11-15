package img

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/usecase"
	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		uc  ImgUsecase
		log *logrus.Logger
	}

	ImgUsecase interface {
		GetAllLogos(ctx context.Context) (entity.LogosResp, error)
		Get(ctx context.Context, params usecase.PageParams) (*entity.ImagesResp, error)
		Create(ctx context.Context, image *entity.Image) (*entity.Image, error)
		Update(ctx context.Context, id int, image *entity.ImageUpdate) (*entity.Image, error)
		Delete(ctx context.Context, id int) error
	}
)

func New(uc ImgUsecase, log *logrus.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

type (
	ImgUC struct {
		db    DBXT
		repo  ImgRepo
		cache cache.Cache
		log   *logrus.Logger
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	ImgRepo interface {
		GetAllLogos(ctx context.Context, qe repository.Querier) (entity.LogosResp, error)
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Image, error)
		Create(ctx context.Context, qe repository.Querier, image *entity.Image) (*entity.Image, error)
		Update(ctx context.Context, qe repository.Querier, id int, image *entity.ImageUpdate) (*entity.Image, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}
)

func New(db DBXT, repo ImgRepo, cache cache.Cache, log *logrus.Logger) *ImgUC {
	return &ImgUC{
		db: db, 
		repo: repo,
		cache: cache,}
}

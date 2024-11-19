package adUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
)

type (
	AdUC struct {
		db    DBXT
		repo  AdRepo
		cache cache.Cache
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	AdRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.AdvResp, error)
		GetWithFilter(ctx context.Context, qe repository.Querier, params repository.AdvFilterParams) ([]*entity.AdvResp, error) 
		Create(ctx context.Context, qe repository.Querier, adv *entity.Adv) (*entity.AdvResp, error)
		Update(ctx context.Context, qe repository.Querier, id int, adv *entity.AdvUpdate) (*entity.AdvResp, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}
)

func New(db DBXT, repo AdRepo, cache cache.Cache) *AdUC {
	return &AdUC{
		db:    db,
		repo:  repo,
		cache: cache,
	}
}

package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/cache/errcache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

const (
	logoCachePrefix  = "logo:*"
	allLogosCachekey = "logo:all"
)

type LogoUsecase struct {
	repo  repository.LogoRepo
	cache cache.Cache
	log   *logrus.Logger
}

func NewLogoUsecase(repo repository.LogoRepo, cache cache.Cache, log *logrus.Logger) *LogoUsecase {
	return &LogoUsecase{repo: repo, cache: cache, log: log}
}

func (uc *LogoUsecase) DeleteCache(ctx context.Context) {
	if err := uc.cache.Delete(context.Background(), logoCachePrefix); err != nil {
		uc.log.Error(err)
	}
}

func (uc *LogoUsecase) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	logos, err := uc.repo.Add(ctx, logo)
	if err != nil {
		return nil, err
	}

	uc.DeleteCache(context.Background())
	return logos, err
}

func (uc *LogoUsecase) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	logos, err := uc.repo.Update(ctx, id, logo)
	if err != nil {
		return nil, err
	}

	uc.DeleteCache(context.Background())
	return logos, err
}

func (uc *LogoUsecase) Delete(ctx context.Context, id int) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}

	uc.DeleteCache(context.Background())
	return nil
}

func (uc *LogoUsecase) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *LogoUsecase) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	var logos []*entity.LogoResp

	logosStr, err := uc.cache.Get(ctx, allLogosCachekey)
	if err != nil {
		if !errors.Is(err, errcache.ErrKeyNotFound) {
			uc.log.Error(err)
		}

		logos, err = uc.repo.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(ctx, allLogosCachekey, logos); err != nil {
			uc.log.Error(err)
		}

		return logos, nil
	}

	err = json.Unmarshal([]byte(logosStr), &logos)
	return logos, err
}

func (uc *LogoUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.LogosResp, error) {
	logos, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.LogosResp{
		Logos: logos,
		Total: count,
	}, err
}

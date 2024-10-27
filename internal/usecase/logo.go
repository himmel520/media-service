package usecase

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type LogoUsecase struct {
	repo repository.LogoRepo
	log  *logrus.Logger
}

func NewLogoUsecase(repo repository.LogoRepo, log *logrus.Logger) *LogoUsecase {
	return &LogoUsecase{repo: repo, log: log}
}

func (uc *LogoUsecase) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	return uc.repo.Add(ctx, logo)
}

func (uc *LogoUsecase) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	return uc.repo.Update(ctx, id, logo)
}

func (uc *LogoUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *LogoUsecase) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *LogoUsecase) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	return uc.repo.GetAll(ctx)
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

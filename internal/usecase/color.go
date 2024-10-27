package usecase

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type ColorUsecase struct {
	repo repository.ColorRepo
	log  *logrus.Logger
}

func NewColorUsecase(repo repository.ColorRepo, log *logrus.Logger) *ColorUsecase {
	return &ColorUsecase{repo: repo, log: log}
}

func (uc *ColorUsecase) Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error) {
	return uc.repo.Add(ctx, color)
}

func (uc *ColorUsecase) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.ColorResp, error) {
	return uc.repo.Update(ctx, id, color)
}

func (uc *ColorUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *ColorUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.ColorsResp, error) {
	colors, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ColorsResp{
		Colors: colors,
		Total:  count,
	}, err
}

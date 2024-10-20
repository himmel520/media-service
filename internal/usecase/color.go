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

func (s *ColorUsecase) Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error) {
	return s.repo.Add(ctx, color)
}

func (s *ColorUsecase) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.ColorResp, error) {
	return s.repo.Update(ctx, id, color)
}

func (s *ColorUsecase) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *ColorUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.ColorsResp, error) {
	colors, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ColorsResp{
		Colors: colors,
		Total:  count,
	}, err
}

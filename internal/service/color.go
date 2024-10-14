package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/models"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --all

type ColorRepo interface {
	Add(ctx context.Context, color *models.Color) (*models.ColorResp, error)
	Update(ctx context.Context, id int, Color *models.ColorUpdate) (*models.ColorResp, error)
	Delete(ctx context.Context, id int) error
	GetAllWithPagination(ctx context.Context, limit, offset int) ([]*models.ColorResp, error)
	Count(ctx context.Context) (int, error)
}

type ColorService struct {
	repo ColorRepo
	log  *logrus.Logger
}

func NewColorService(repo ColorRepo, log *logrus.Logger) *ColorService {
	return &ColorService{repo: repo, log: log}
}

func (s *ColorService) Add(ctx context.Context, color *models.Color) (*models.ColorResp, error) {
	return s.repo.Add(ctx, color)
}

func (s *ColorService) Update(ctx context.Context, id int, color *models.ColorUpdate) (*models.ColorResp, error) {
	return s.repo.Update(ctx, id, color)
}

func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *ColorService) GetAllWithPagination(ctx context.Context, limit, offset int) (*models.ColorsResp, error) {
	colors, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &models.ColorsResp{
		Colors: colors,
		Total:  count,
	}, err
}

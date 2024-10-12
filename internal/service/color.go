package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/models"
)

func (s *Service) AddColor(ctx context.Context, color *models.Color) (*models.ColorResp, error) {
	return s.repo.AddColor(ctx, color)
}

func (s *Service) UpdateColor(ctx context.Context, id int, color *models.ColorUpdate) (*models.ColorResp, error) {
	return s.repo.UpdateColor(ctx, id, color)
}

func (s *Service) DeleteColor(ctx context.Context, id int) error {
	return s.repo.DeleteColor(ctx, id)
}

func (s *Service) GetColors(ctx context.Context, limit, offset int) (*models.ColorsResp, error) {
	colors, err := s.repo.GetColors(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.GetColorCount(ctx)
	if err != nil {
		return nil, err
	}

	return &models.ColorsResp{
		Colors: colors,
		Total:  count,
	}, err
}

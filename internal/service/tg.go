package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/models"
)

func (s *Service) AddTG(ctx context.Context, tg *models.TG) (*models.TGResp, error) {
	return s.repo.AddTG(ctx, tg)
}

func (s *Service) UpdateTG(ctx context.Context, id int, TG *models.TGUpdate) (*models.TGResp, error) {
	return s.repo.UpdateTG(ctx, id, TG)
}

func (s *Service) DeleteTG(ctx context.Context, id int) error {
	return s.repo.DeleteTG(ctx, id)
}

func (s *Service) GetTGs(ctx context.Context, limit, offset int) (*models.TGsResp, error) {
	tgs, err := s.repo.GetTGs(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.GetTGCount(ctx)
	if err != nil {
		return nil, err
	}

	return &models.TGsResp{
		TGs:   tgs,
		Total: count,
	}, err
}

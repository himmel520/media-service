package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/models"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --all

type TGRepo interface {
	Add(ctx context.Context, tg *models.TG) (*models.TGResp, error)
	Update(ctx context.Context, id int, tg *models.TGUpdate) (*models.TGResp, error)
	Delete(ctx context.Context, id int) error
	GetAllWithPagination(ctx context.Context, limit, offset int) ([]*models.TGResp, error)
	Count(ctx context.Context) (int, error)
}

type TGService struct {
	repo TGRepo
	log  *logrus.Logger
}

func NewTGService(repo TGRepo, log *logrus.Logger) *TGService {
	return &TGService{repo: repo, log: log}
}

func (s *TGService) Add(ctx context.Context, tg *models.TG) (*models.TGResp, error) {
	return s.repo.Add(ctx, tg)
}

func (s *TGService) Update(ctx context.Context, id int, TG *models.TGUpdate) (*models.TGResp, error) {
	return s.repo.Update(ctx, id, TG)
}

func (s *TGService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *TGService) GetAllWithPagination(ctx context.Context, limit, offset int) (*models.TGsResp, error) {
	tgs, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &models.TGsResp{
		TGs:   tgs,
		Total: count,
	}, err
}

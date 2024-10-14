package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/models"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --all

type LogoRepo interface {
	Add(ctx context.Context, logo *models.Logo) (*models.LogoResp, error)
	Update(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, logoID int) (*models.LogoResp, error)
	GetAllWithPagination(ctx context.Context, limit, offset int) (map[int]*models.Logo, error)
	GetAll(ctx context.Context) ([]*models.Logo, error)
	Count(ctx context.Context) (int, error)
}

type LogoService struct {
	repo LogoRepo
	log  *logrus.Logger
}

func NewLogoService(repo LogoRepo, log *logrus.Logger) *LogoService {
	return &LogoService{repo: repo, log: log}
}

func (s *LogoService) Add(ctx context.Context, logo *models.Logo) (*models.LogoResp, error) {
	return s.repo.Add(ctx, logo)
}

func (s *LogoService) Update(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error) {
	return s.repo.Update(ctx, id, logo)
}

func (s *LogoService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *LogoService) GetByID(ctx context.Context, id int) (*models.LogoResp, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LogoService) GetAll(ctx context.Context) ([]*models.Logo, error) {
	return s.repo.GetAll(ctx)
}

func (s *LogoService) GetAllWithPagination(ctx context.Context, limit, offset int) (*models.LogosResp, error) {
	logos, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &models.LogosResp{
		Logos: logos,
		Total: count,
	}, err
}

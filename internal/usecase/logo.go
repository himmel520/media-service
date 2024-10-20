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

func (s *LogoUsecase) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	return s.repo.Add(ctx, logo)
}

func (s *LogoUsecase) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	return s.repo.Update(ctx, id, logo)
}

func (s *LogoUsecase) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *LogoUsecase) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LogoUsecase) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	return s.repo.GetAll(ctx)
}

func (s *LogoUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.LogosResp, error) {
	logos, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.LogosResp{
		Logos: logos,
		Total: count,
	}, err
}

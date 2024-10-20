package usecase

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/internal/entity"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type TGUsecase struct {
	repo repository.TGRepo
	log  *logrus.Logger
}

func NewTGUsecase(repo repository.TGRepo, log *logrus.Logger) *TGUsecase {
	return &TGUsecase{repo: repo, log: log}
}

func (s *TGUsecase) Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error) {
	return s.repo.Add(ctx, tg)
}

func (s *TGUsecase) Update(ctx context.Context, id int, TG *entity.TGUpdate) (*entity.TGResp, error) {
	return s.repo.Update(ctx, id, TG)
}

func (s *TGUsecase) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *TGUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.TGsResp, error) {
	tgs, err := s.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.TGsResp{
		TGs:   tgs,
		Total: count,
	}, err
}

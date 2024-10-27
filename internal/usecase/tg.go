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

func (uc *TGUsecase) Add(ctx context.Context, tg *entity.TG) (*entity.TGResp, error) {
	return uc.repo.Add(ctx, tg)
}

func (uc *TGUsecase) Update(ctx context.Context, id int, TG *entity.TGUpdate) (*entity.TGResp, error) {
	return uc.repo.Update(ctx, id, TG)
}

func (uc *TGUsecase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *TGUsecase) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.TGsResp, error) {
	tgs, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.TGsResp{
		TGs:   tgs,
		Total: count,
	}, err
}

package service

import (
	"context"

	"github.com/himmel520/uoffer/mediaAd/models"
)

// аналогично с прошлым микросом - можно заменить две команды ниже на одну - set
// для того что бы сделать операцию идемпотентной
func (s *Service) AddLogo(ctx context.Context, logo *models.Logo) (*models.LogoResp, error) {
	return s.repo.AddLogo(ctx, logo)
}

func (s *Service) UpdateLogo(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error) {
	return s.repo.UpdateLogo(ctx, id, logo)
}

func (s *Service) DeleteLogo(ctx context.Context, id int) error {
	return s.repo.DeleteLogo(ctx, id)
}

func (s *Service) GetLogo(ctx context.Context, id int) (*models.LogoResp, error) {
	return s.repo.GetLogo(ctx, id)
}

func (s *Service) GetLogos(ctx context.Context, limit, offset int) (*models.LogosResp, error) {
	logos, err := s.repo.GetLogos(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.GetLogoCount(ctx)
	if err != nil {
		return nil, err
	}

	return &models.LogosResp{
		Logos: logos,
		Total: count,
	}, err
}

package imgUC

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache/errcache"
)

func (uc *ImgUC) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ImgUC) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	var logos []*entity.LogoResp

	logosStr, err := uc.cache.Get(ctx, allLogosCachekey)
	if err != nil {
		if !errors.Is(err, errcache.ErrKeyNotFound) {
			uc.log.Error(err)
		}

		logos, err = uc.repo.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(ctx, allLogosCachekey, logos); err != nil {
			uc.log.Error(err)
		}

		return logos, nil
	}

	err = json.Unmarshal([]byte(logosStr), &logos)
	return logos, err
}

func (uc *ImgUC) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.LogosResp, error) {
	logos, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.LogosResp{
		Logos: logos,
		Total: count,
	}, err
}

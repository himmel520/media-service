package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
)

func (uc *ImgUC) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
	logos, err := uc.repo.Update(ctx, id, logo)
	if err != nil {
		return nil, err
	}

	uc.DeleteCache(context.Background())
	return logos, err
}

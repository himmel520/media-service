package tgRepo

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *TgRepo) GetAllWithPagination(ctx context.Context, qe repository.Querier, limit, offset int) ([]*entity.TGResp, error) {
	rows, err := qe.Query(ctx, `
	select * 
		from tg
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tgs := []*entity.TGResp{}
	for rows.Next() {
		tg := &entity.TGResp{}
		if err := rows.Scan(&tg.ID, &tg.Title, &tg.Url); err != nil {
			return nil, err
		}

		tgs = append(tgs, tg)
	}

	if len(tgs) == 0 {
		return nil, repoerr.ErrTGNotFound
	}

	return tgs, err
}


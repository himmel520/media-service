package colorRepo

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *ColorRepo) GetAllWithPagination(ctx context.Context, qe repository.Querier, limit, offset int) ([]*entity.ColorResp, error) {
	rows, err := qe.Query(ctx, `
	select * 
		from colors
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []*entity.ColorResp{}
	for rows.Next() {
		color := &entity.ColorResp{}
		if err := rows.Scan(&color.ID, &color.Title, &color.Hex); err != nil {
			return nil, err
		}

		colors = append(colors, color)
	}

	if len(colors) == 0 {
		return nil, repoerr.ErrColorNotFound
	}

	return colors, err
}

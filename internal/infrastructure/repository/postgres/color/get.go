package colorRepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *ColorRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Color, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"title",
			"hex").
		From("colors").
		OrderBy("title").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []*entity.Color{}
	for rows.Next() {
		color := &entity.Color{}
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

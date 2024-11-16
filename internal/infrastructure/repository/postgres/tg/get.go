package tgRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *TgRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Tg, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"title",
			"url").
		From("tg").
		OrderBy("title").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	tgs := []*entity.Tg{}
	for rows.Next() {
		tg := &entity.Tg{}
		if err := rows.Scan(
			&tg.ID,
			&tg.Title,
			&tg.Url); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		tgs = append(tgs, tg)
	}

	if len(tgs) == 0 {
		return nil, repoerr.ErrTGNotFound
	}

	return tgs, err
}

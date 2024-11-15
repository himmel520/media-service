package colorRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ColorRepo) Create(ctx context.Context, qe repository.Querier, color *entity.Color) (*entity.Color, error) {
	query, args, err := squirrel.Insert("colors").
		Columns(
			"title",
			"hex").
		Values(
			color.Title,
			color.Hex).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id, title, hex").
		ToSql()
	if err != nil {
		return nil, err
	}

	newColor := &entity.Color{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newColor.ID,
		&newColor.Title,
		&newColor.Hex)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}

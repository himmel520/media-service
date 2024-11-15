package colorRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ColorRepo) Update(ctx context.Context, qe repository.Querier, id int, color *entity.ColorUpdate) (*entity.Color, error) {
	builder := squirrel.Update("colors").
		Where(squirrel.Eq{"id": id}).
		Suffix(`
		returning 
			id, 
			title, 
			hex`).
		PlaceholderFormat(squirrel.Dollar)

	if color.Title.Set {
		builder = builder.Set("title", color.Title.Value)

	}

	if color.Hex.Set {
		builder = builder.Set("hex", color.Hex.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	newColor := &entity.Color{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newColor.ID,
		&newColor.Title,
		&newColor.Hex)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}

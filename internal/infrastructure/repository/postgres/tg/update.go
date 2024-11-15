package tgRepo

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

func (r *TgRepo) Update(ctx context.Context, qe repository.Querier, id int, tg *entity.TgUpdate) (*entity.Tg, error) {
	builder := squirrel.Update("tg").
		Where(squirrel.Eq{"id": id}).
		Suffix(`
		returning 
			id, 
			title, 
			url`).
		PlaceholderFormat(squirrel.Dollar)

	if tg.Url.Set {
		builder = builder.Set("url", tg.Url.Value)
	}

	if tg.Title.Set {
		builder = builder.Set("title", tg.Title.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	newTG := &entity.Tg{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newTG.ID,
		&newTG.Title,
		&newTG.Url)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrTGNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrTGExist
	}

	return newTG, err
}

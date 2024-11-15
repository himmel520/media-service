package tgRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *TgRepo) Create(ctx context.Context, qe repository.Querier, tg *entity.Tg) (*entity.Tg, error) {
	qyery, args, err := squirrel.Insert("tg").
		Columns(
			"title",
			"url").
		Values(
			tg.Title,
			tg.Url).
		Suffix("returning id, title, url").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	newTG := &entity.Tg{}
	err = qe.QueryRow(ctx, qyery, args...).Scan(
		&newTG.ID, 
		&newTG.Title, 
		&newTG.Url)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrTGExist
	}

	return newTG, err
}

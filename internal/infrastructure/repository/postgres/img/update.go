package imgRepo

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

func (r *ImgRepo) Update(ctx context.Context, qe repository.Querier, id int, image *entity.ImageUpdate) (*entity.Image, error) {
	builder := squirrel.Update("images").
		Where(squirrel.Eq{"id": id}).
		Suffix(`
		returning 
		id, 
		url,
		title, 
		type`).
		PlaceholderFormat(squirrel.Dollar)

	if image.Url.Set {
		builder = builder.Set("url", image.Url.Value)
	}

	if image.Title.Set {
		builder = builder.Set("title", image.Title.Value)
	}

	if image.Type.Set {
		builder = builder.Set("type", image.Type.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	newLogo := &entity.Image{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newLogo.ID,
		&newLogo.Url,
		&newLogo.Title,
		&newLogo.Type)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrImageNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrImageExist
	}

	return newLogo, err
}

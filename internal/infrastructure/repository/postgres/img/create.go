package imgRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ImgRepo) Create(ctx context.Context, qe repository.Querier, image *entity.Image) (*entity.Image, error) {
	query, args, err := squirrel.Insert("images").
		Columns(
			"url",
			"title",
			"type").
		Values(
			image.Url,
			image.Title,
			image.Type).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(`
		returning 
			id, 
			url,
			title, 
			type`).
		ToSql()
	if err != nil {
		return nil, err
	}

	newImage := &entity.Image{}

	err = qe.QueryRow(ctx, query, args...).Scan(
		&newImage.ID, 
		&newImage.Url,
		&newImage.Title,
		&newImage.Type)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrImageExist
	}

	return newImage, err
}

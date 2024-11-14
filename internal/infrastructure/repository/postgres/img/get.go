package imgRepo

import (
	"context"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
)

func (r *ImgRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Image, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"url",
			"title",
			"type").
		From("images").
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

	images := []*entity.Image{}
	for rows.Next() {
		image := &entity.Image{}
		if err := rows.Scan(
			&image.ID,
			&image.Url,
			&image.Title,
			&image.Type); err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	if len(images) == 0 {
		return nil, repoerr.ErrImageNotFound
	}

	return images, err
}

func (r *ImgRepo) GetAllLogos(ctx context.Context, qe repository.Querier) (entity.LogosResp, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"url",
			"title",
			"type").
		From("images").
		OrderBy("title").Where(squirrel.Eq{"type": "logo"}).
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

	images := entity.LogosResp{}
	for rows.Next() {
		image := &entity.Image{}
		if err := rows.Scan(
			&image.ID,
			&image.Url,
			&image.Title,
			&image.Type); err != nil {
			return nil, err
		}

		images[strconv.Itoa(image.ID)] = image
	}

	if len(images) == 0 {
		return nil, repoerr.ErrImageNotFound
	}

	return images, err
}

package adRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
)

func (r *AdRepo) GetByID(ctx context.Context, qe repository.Querier, id int) (*entity.AdvResp, error) {
	query, args, err := squirrel.Select(
		"i.id",
		"i.url",
		"i.title",
		"i.type",
		"t.id",
		"t.title",
		"t.url",
		"c.id",
		"c.title",
		"c.hex",
		"a.id",
		"a.post",
		"a.title",
		"a.description",
		"a.priority").
		From("adv AS a").
		Join("colors AS c ON c.id = a.colors_id").
		Join("tg AS t ON t.id = a.tg_id").
		Join("images AS i ON i.id = a.images_id").
		Where(squirrel.Eq{"a.id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	advResp := &entity.AdvResp{}
	if err = qe.QueryRow(ctx, query, args...).Scan(
		&advResp.Image.ID,
		&advResp.Image.Url,
		&advResp.Image.Title,
		&advResp.Image.Type,
		&advResp.Tg.ID,
		&advResp.Tg.Title,
		&advResp.Tg.Url,
		&advResp.Color.ID,
		&advResp.Color.Title,
		&advResp.Color.Hex,
		&advResp.ID,
		&advResp.Post,
		&advResp.Title,
		&advResp.Description,
		&advResp.Priority); err != nil {
		return nil, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrAdvNotFound
	}

	return advResp, err
}

func (r *AdRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.AdvResp, error) {
	query, args, err := squirrel.Select(
		"i.id",
		"i.url",
		"i.title",
		"i.type",
		"t.id",
		"t.title",
		"t.url",
		"c.id",
		"c.title",
		"c.hex",
		"a.id",
		"a.post",
		"a.title",
		"a.description",
		"a.priority").
		From("adv AS a").
		Join("colors AS c ON c.id = a.colors_id").
		Join("tg AS t ON t.id = a.tg_id").
		Join("images AS i ON i.id = a.images_id").
		OrderBy("a.post").
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

	advResp := []*entity.AdvResp{}
	for rows.Next() {
		adv := &entity.AdvResp{}
		if err = rows.Scan(
			&adv.Image.ID,
			&adv.Image.Url,
			&adv.Image.Title,
			&adv.Image.Type,
			&adv.Tg.ID,
			&adv.Tg.Title,
			&adv.Tg.Url,
			&adv.Color.ID,
			&adv.Color.Title,
			&adv.Color.Hex,
			&adv.ID,
			&adv.Post,
			&adv.Title,
			&adv.Description,
			&adv.Priority); err != nil {
			return nil, err
		}

		advResp = append(advResp, adv)
	}

	if len(advResp) == 0 {
		return nil, repoerr.ErrAdvNotFound
	}

	return advResp, err
}

func (r *AdRepo) GetWithFilter(ctx context.Context, qe repository.Querier, params repository.AdvFilterParams) ([]*entity.AdvResp, error) {
	builder := squirrel.Select(
		"i.id",
		"i.url",
		"i.title",
		"i.type",
		"t.id",
		"t.title",
		"t.url",
		"c.id",
		"c.title",
		"c.hex",
		"a.id",
		"a.post",
		"a.title",
		"a.description",
		"a.priority").
		From("adv AS a").
		Join("colors AS c ON c.id = a.colors_id").
		Join("tg AS t ON t.id = a.tg_id").
		Join("images AS i ON i.id = a.images_id").
		OrderBy("a.post").
		PlaceholderFormat(squirrel.Dollar)

	if len(params.Posts) > 0 {
		builder = builder.Where(squirrel.Eq{"a.post": params.Posts})
	}

	if len(params.Priority) > 0 {
		builder = builder.Where(squirrel.Eq{"a.priority": params.Priority})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	advResp := []*entity.AdvResp{}
	for rows.Next() {
		adv := &entity.AdvResp{}
		if err = rows.Scan(
			&adv.Image.ID,
			&adv.Image.Url,
			&adv.Image.Title,
			&adv.Image.Type,
			&adv.Tg.ID,
			&adv.Tg.Title,
			&adv.Tg.Url,
			&adv.Color.ID,
			&adv.Color.Title,
			&adv.Color.Hex,
			&adv.ID,
			&adv.Post,
			&adv.Title,
			&adv.Description,
			&adv.Priority); err != nil {
			return nil, err
		}

		advResp = append(advResp, adv)
	}

	if len(advResp) == 0 {
		return nil, repoerr.ErrAdvNotFound
	}

	return advResp, err
}

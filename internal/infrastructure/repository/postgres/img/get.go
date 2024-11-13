package imgRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
)

func (r *ImgRepo) GetByID(ctx context.Context, qe repository.Querier, id int) (*entity.LogoResp, error) {
	logo := &entity.LogoResp{}

	err := qe.QueryRow(ctx, `
	select * from logos 
		where id = $1;`, id).Scan(&logo.ID, &logo.Url, &logo.Title)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrLogoNotFound
	}

	return logo, err
}

func (r *ImgRepo) GetAll(ctx context.Context, qe repository.Querier) ([]*entity.LogoResp, error) {
	rows, err := qe.Query(ctx, `
	select * 
		from logos
	order by title asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logos := []*entity.LogoResp{}
	for rows.Next() {
		logo := &entity.LogoResp{}
		if err := rows.Scan(&logo.ID, &logo.Url, &logo.Title); err != nil {
			return nil, err
		}

		logos = append(logos, logo)
	}

	if len(logos) == 0 {
		return nil, repoerr.ErrLogoNotFound
	}

	return logos, err
}

func (r *ImgRepo) GetAllWithPagination(ctx context.Context, qe repository.Querier, limit, offset int) (map[int]*entity.Logo, error) {
	rows, err := qe.Query(ctx, `
	select * 
		from logos
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logos := map[int]*entity.Logo{}
	for rows.Next() {
		logo := &entity.Logo{}
		if err := rows.Scan(&logo.ID, &logo.Url, &logo.Title); err != nil {
			return nil, err
		}

		logos[logo.ID] = logo
	}

	if len(logos) == 0 {
		return nil, repoerr.ErrLogoNotFound
	}

	return logos, err
}
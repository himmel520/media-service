package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/repository"
	"github.com/himmel520/uoffer/mediaAd/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// получается так что у тебя имея доступ к *Repository можно любые запросы к базе сделать для любой сущности
// в идеале что то типа LogoRepository
func (r *Repository) AddLogo(ctx context.Context, logo *models.Logo) (*models.LogoResp, error) {
	newLogo := &models.LogoResp{}

	err := r.DB.QueryRow(ctx, `
	insert into logos
		(url, title)
	values
		($1, $2)
	returning *`, logo.Url, logo.Title).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint {
		return nil, repository.ErrLogoExist
	}

	return newLogo, err
}

func (r *Repository) UpdateLogo(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error) {
	var keys []string
	var values []interface{}
	if logo.Url != nil {
		keys = append(keys, "url=$1")
		values = append(values, logo.Url)
	}

	if logo.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(values)+1))
		values = append(values, logo.Title)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update logos
		set %v
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newLogo := &models.LogoResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repository.ErrLogoNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint:
		return nil, repository.ErrLogoExist
	}

	return newLogo, err
}

func (r *Repository) DeleteLogo(ctx context.Context, id int) error {
	var pgErr *pgconn.PgError

	cmdTag, err := r.DB.Exec(ctx, `
	delete from logos
		where id = $1`, id)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.FKViolation {
			return repository.ErrLogoDependency
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repository.ErrLogoNotFound
	}

	return err
}

func (r *Repository) GetLogo(ctx context.Context, id int) (*models.LogoResp, error) {
	logo := &models.LogoResp{}

	err := r.DB.QueryRow(ctx, `
	select * from logos
		where id = $1;`, id).Scan(&logo.ID, &logo.Url, &logo.Title)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrLogoNotFound
	}

	return logo, err
}

func (r *Repository) GetLogos(ctx context.Context, limit, offset int) (map[int]*models.Logo, error) {
	rows, err := r.DB.Query(ctx, `
	select *
		from logos
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logos := map[int]*models.Logo{}
	for rows.Next() {
		logo := &models.Logo{}
		if err := rows.Scan(&logo.ID, &logo.Url, &logo.Title); err != nil {
			return nil, err
		}

		logos[logo.ID] = logo
	}

	if len(logos) == 0 {
		return nil, repository.ErrLogoNotFound
	}

	return logos, err
}

func (r *Repository) GetLogoCount(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM logos;`).Scan(&count)
	return count, err
}

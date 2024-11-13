package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LogoRepo struct {
	DB *pgxpool.Pool
}

func NewLogorRepo(db *pgxpool.Pool) *LogoRepo {
	return &LogoRepo{DB: db}
}

func (r *LogoRepo) Add(ctx context.Context, logo *entity.Logo) (*entity.LogoResp, error) {
	newLogo := &entity.LogoResp{}

	err := r.DB.QueryRow(ctx, `
	insert into logos 
		(url, title) 
	values 
		($1, $2) 
	returning *`, logo.Url, logo.Title).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrLogoExist
	}

	return newLogo, err
}

func (r *LogoRepo) Update(ctx context.Context, id int, logo *entity.LogoUpdate) (*entity.LogoResp, error) {
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

	newLogo := &entity.LogoResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrLogoNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrLogoExist
	}

	return newLogo, err
}

func (r *LogoRepo) Delete(ctx context.Context, id int) error {
	var pgErr *pgconn.PgError

	cmdTag, err := r.DB.Exec(ctx, `
	delete from logos 
		where id = $1`, id)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return repoerr.ErrLogoDependency
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrLogoNotFound
	}

	return err
}

func (r *LogoRepo) GetByID(ctx context.Context, id int) (*entity.LogoResp, error) {
	logo := &entity.LogoResp{}

	err := r.DB.QueryRow(ctx, `
	select * from logos 
		where id = $1;`, id).Scan(&logo.ID, &logo.Url, &logo.Title)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrLogoNotFound
	}

	return logo, err
}

func (r *LogoRepo) GetAll(ctx context.Context) ([]*entity.LogoResp, error) {
	rows, err := r.DB.Query(ctx, `
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

func (r *LogoRepo) GetAllWithPagination(ctx context.Context, limit, offset int) (map[int]*entity.Logo, error) {
	rows, err := r.DB.Query(ctx, `
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

func (r *LogoRepo) Count(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM logos;`).Scan(&count)
	return count, err
}

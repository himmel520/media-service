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

func (r *Repository) AddColor(ctx context.Context, color *models.Color) (*models.ColorResp, error) {
	newColor := &models.ColorResp{}

	err := r.DB.QueryRow(ctx, `
	insert into colors 
		(title, hex) 
	values 
		($1, $2) 
	returning *`, color.Title, color.Hex).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint {
		return nil, repository.ErrColorHexExist
	}

	return newColor, err
}

func (r *Repository) UpdateColor(ctx context.Context, id int, color *models.ColorUpdate) (*models.ColorResp, error) {
	var keys []string
	var values []interface{}
	if color.Title != nil {
		keys = append(keys, "title=$1")
		values = append(values, color.Title)
	}

	if color.Hex != nil {
		keys = append(keys, fmt.Sprintf("hex=$%d", len(values)+1))
		values = append(values, color.Hex)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update colors 
		set %v 
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newColor := &models.ColorResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint {
		return nil, repository.ErrColorHexExist
	}

	return newColor, err
}

func (r *Repository) DeleteColor(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `
	delete from Colors 
		where id = $1`, id)

	if cmdTag.RowsAffected() == 0 {
		return repository.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repository.FKViolation {
		return repository.ErrColorDependencyExist
	}


	return err
}

func (r *Repository) GetColors(ctx context.Context, limit, offset int) ([]*models.ColorResp, error) {
	rows, err := r.DB.Query(ctx, `
	select * 
		from colors
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []*models.ColorResp{}
	for rows.Next() {
		color := &models.ColorResp{}
		if err := rows.Scan(&color.ID, &color.Title, &color.Hex); err != nil {
			return nil, err
		}

		colors = append(colors, color)
	}

	if len(colors) == 0 {
		return nil, repository.ErrColorNotFound
	}

	return colors, err
}

func (r *Repository) GetColorCount(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM colors;`).Scan(&count)
	return count, err
}

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

type ColorRepo struct {
	DB *pgxpool.Pool
}

func NewColorRepo(db *pgxpool.Pool) *ColorRepo {
	return &ColorRepo{DB: db}
}

func (r *ColorRepo) Add(ctx context.Context, color *entity.Color) (*entity.ColorResp, error) {
	newColor := &entity.ColorResp{}

	err := r.DB.QueryRow(ctx, `
	insert into colors 
		(title, hex) 
	values 
		($1, $2) 
	returning *`, color.Title, color.Hex).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}

func (r *ColorRepo) Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.ColorResp, error) {
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

	newColor := &entity.ColorResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newColor.ID, &newColor.Title, &newColor.Hex)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrColorHexExist
	}

	return newColor, err
}

func (r *ColorRepo) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `
	delete from Colors 
		where id = $1`, id)

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrColorNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.FKViolation {
		return repoerr.ErrColorDependencyExist
	}

	return err
}

func (r *ColorRepo) GetAllWithPagination(ctx context.Context, limit, offset int) ([]*entity.ColorResp, error) {
	rows, err := r.DB.Query(ctx, `
	select * 
		from colors
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []*entity.ColorResp{}
	for rows.Next() {
		color := &entity.ColorResp{}
		if err := rows.Scan(&color.ID, &color.Title, &color.Hex); err != nil {
			return nil, err
		}

		colors = append(colors, color)
	}

	if len(colors) == 0 {
		return nil, repoerr.ErrColorNotFound
	}

	return colors, err
}

func (r *ColorRepo) Count(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM colors;`).Scan(&count)
	return count, err
}

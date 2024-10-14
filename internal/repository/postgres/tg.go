package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/mediaAd/internal/models"
	"github.com/himmel520/uoffer/mediaAd/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TGRepo struct {
	DB *pgxpool.Pool
}

func NewTGRepo(db *pgxpool.Pool) *TGRepo {
	return &TGRepo{DB: db}
}

func (r *TGRepo) Add(ctx context.Context, tg *models.TG) (*models.TGResp, error) {
	newTG := &models.TGResp{}

	err := r.DB.QueryRow(ctx, `
	insert into tg 
		(url, title) 
	values 
		($1, $2) 
	returning *`, tg.Url, tg.Title).Scan(&newTG.ID, &newTG.Title, &newTG.Url)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint {
		return nil, repository.ErrTGExist
	}

	return newTG, err
}

func (r *TGRepo) Update(ctx context.Context, id int, tg *models.TGUpdate) (*models.TGResp, error) {
	var keys []string
	var values []interface{}
	if tg.Url != nil {
		keys = append(keys, "url=$1")
		values = append(values, tg.Url)
	}

	if tg.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(values)+1))
		values = append(values, tg.Title)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update tg 
		set %v 
	where id = $%v
	returning *;`, strings.Join(keys, ", "), len(values))

	newTG := &models.TGResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newTG.ID, &newTG.Title, &newTG.Url)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repository.ErrTGNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repository.UniqueConstraint:
		return nil, repository.ErrTGExist
	}

	return newTG, err
}

func (r *TGRepo) Delete(ctx context.Context, id int) error {
	var pgErr *pgconn.PgError

	cmdTag, err := r.DB.Exec(ctx, `
	delete from TGs 
		where id = $1`, id)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.FKViolation {
			return repository.ErrTGDependencyExist
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repository.ErrTGNotFound
	}

	return err
}

func (r *TGRepo) GetAllWithPagination(ctx context.Context, limit, offset int) ([]*models.TGResp, error) {
	rows, err := r.DB.Query(ctx, `
	select * 
		from tg
	order by title asc
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tgs := []*models.TGResp{}
	for rows.Next() {
		tg := &models.TGResp{}
		if err := rows.Scan(&tg.ID, &tg.Title, &tg.Url); err != nil {
			return nil, err
		}

		tgs = append(tgs, tg)
	}

	if len(tgs) == 0 {
		return nil, repository.ErrTGNotFound
	}

	return tgs, err
}

func (r *TGRepo) Count(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM tg;`).Scan(&count)
	return count, err
}

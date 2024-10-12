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

// Adv interface {
// 	// Переделать
// 	// AddAdv(ctx context.Context, adv *models.Adv) (int, error)
// 	// DeleteAdv(ctx context.Context, id int) error
// 	// UpdateAdv(ctx context.Context, id int, adv *models.AdvUpdate) error
// 	// Получить все adv по должности
// 	// получить все adv по priority и posts
// }

func (r *Repository) AddAdv(ctx context.Context, adv *models.Adv) (int, error) {
	var id int
	err := r.DB.QueryRow(ctx, `
	insert into adv
		(logos_id, colors_id, tg_id, post, title, description, priority)
	values
		($1, $2, $3, $4, $5, $6, $7)
	returning id;`,
		adv.LogoID, adv.ColorID, adv.TgID, adv.Post, adv.Title, adv.Description, adv.Priority).Scan(&id)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.FKViolation {
			return 0, repository.ErrAdvDependencyNotExist
		}
	}

	return id, err
}

func (r *Repository) GetAdvByID(ctx context.Context, id int) (*models.AdvResponse, error) {
	adv := &models.AdvResponse{}
	err := r.DB.QueryRow(ctx, `
	SELECT
		adv.id,
		colors.hex AS color,
		logos.url AS logo_url,
		tg.url AS tg_url,
		adv.post,
		adv.title,
		adv.description,
		adv.priority
	FROM adv
	JOIN logos ON adv.logos_id = logos.id
	JOIN colors ON adv.colors_id = colors.id
	JOIN tg ON adv.tg_id = tg.id
		WHERE adv.id = $1;`, id).Scan(
		&adv.ID, &adv.Color, &adv.LogoUrl, &adv.TgUrl, &adv.Post,
		&adv.Title, &adv.Description, &adv.Priority)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrAdvNotFound
	}

	return adv, err
}

func (r *Repository) DeleteAdv(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from adv where id = $1`, id)
	if cmdTag.RowsAffected() == 0 {
		return repository.ErrAdvNotFound
	}
	return err

}

func (r *Repository) UpdateAdv(ctx context.Context, id int, adv *models.AdvUpdate) error {
	var keys []string
	var values []interface{}

	if adv.LogoID != nil {
		keys = append(keys, "logos_id=$1")
		values = append(values, adv.LogoID)
	}

	if adv.ColorID != nil {
		keys = append(keys, fmt.Sprintf("colors_id=$%d", len(keys)+1))
		values = append(values, adv.ColorID)
	}

	if adv.TgID != nil {
		keys = append(keys, fmt.Sprintf("tg_id=$%d", len(keys)+1))
		values = append(values, adv.TgID)
	}

	if adv.Post != nil {
		keys = append(keys, fmt.Sprintf("post=$%d", len(keys)+1))
		values = append(values, adv.Post)
	}

	if adv.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(keys)+1))
		values = append(values, adv.Title)
	}

	if adv.Description != nil {
		keys = append(keys, fmt.Sprintf("description=$%d", len(keys)+1))
		values = append(values, adv.Description)
	}

	if adv.Priority != nil {
		keys = append(keys, fmt.Sprintf("priority=$%d", len(keys)+1))
		values = append(values, adv.Priority)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update adv 
		set %s 
	where id = $%d`, strings.Join(keys, ", "), len(values))

	var pgErr *pgconn.PgError
	cmdTag, err := r.DB.Exec(ctx, query, values...)
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.FKViolation {
			return repository.ErrAdvDependencyNotExist
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repository.ErrAdvNotFound
	}

	return err
}

func (r *Repository) GetAdvsWithFilter(ctx context.Context, limit, offset int, posts []string, priority []string) ([]*models.AdvResponse, error) {
	query := `
	SELECT
		adv.id,
		colors.hex AS color,
		logos.url AS logo_url,
		tg.url AS tg_url,
		adv.post,
		adv.title,
		adv.description,
		adv.priority
	FROM adv
	JOIN logos ON adv.logos_id = logos.id
	JOIN colors ON adv.colors_id = colors.id
	JOIN tg ON adv.tg_id = tg.id
	%v
	ORDER BY adv.title ASC
	LIMIT $1 OFFSET $2`

	filter := fmt.Sprintf(`
	WHERE adv.post in ('%v') AND adv.priority in (%v)`, 
	strings.Join(posts, "', '"), strings.Join(priority, ", "))

	rows, err := r.DB.Query(ctx, fmt.Sprintf(query, filter), limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	advs := []*models.AdvResponse{}
	for rows.Next() {
		adv := &models.AdvResponse{}

		if err := rows.Scan(&adv.ID, &adv.Color, &adv.LogoUrl, &adv.TgUrl, &adv.Post,
			&adv.Title, &adv.Description, &adv.Priority); err != nil {
			return nil, err
		}

		advs = append(advs, adv)
	}

	if len(advs) == 0 {
		return nil, repository.ErrAdvNotFound
	}

	return advs, err
}

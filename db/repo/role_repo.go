package repo

import (
	"context"
	"database/sql"
	"errors"

	"tg-star-shop-bot-001/common/domain"
)

type RoleRepo struct {
	db      *sql.DB
	adapter *roleAdapter
}

func NewRoleRepo(db *sql.DB) *RoleRepo {
	return &RoleRepo{
		db:      db,
		adapter: newRoleAdapter(),
	}
}

func (o *RoleRepo) FindByID(ctx context.Context, id int64) (*domain.Role, error) {
	const query = `
        SELECT *
        FROM roles
        WHERE id = $1
    `
	row := o.db.QueryRowContext(ctx, query, id)
	item := &role{}
	if err := row.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt, &item.Code); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(item), nil
}

func (o *RoleRepo) Create(ctx context.Context, item *domain.Role) error {
	const query = `
        INSERT INTO roles (code)
        VALUES ($1)
        RETURNING id
    `
	return o.db.
		QueryRowContext(ctx, query, item.Code).
		Scan(&item.ID)
}

func (o *RoleRepo) Update(ctx context.Context, item *domain.Role) error {
	const query = `
        UPDATE roles
        SET code = $1
        WHERE id = $2
    `
	res, err := o.db.ExecContext(ctx, query, item.Code, item.ID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

func (o *RoleRepo) Delete(ctx context.Context, id int64) error {
	const query = `
        DELETE FROM roles
        WHERE id = $1
    `
	res, err := o.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

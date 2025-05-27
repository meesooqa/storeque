package repo

import (
	"context"
	"database/sql"
	"errors"

	"tg-star-shop-bot-001/common/domain"
)

type UserRepo struct {
	db      *sql.DB
	adapter *userAdapter
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db:      db,
		adapter: newUserAdapter(),
	}
}

func (o *UserRepo) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	const query = `
        SELECT *
        FROM users
        WHERE id = $1
    `
	row := o.db.QueryRowContext(ctx, query, id)
	u := &user{}
	if err := row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.TelegramID, &u.Username, &u.Firstname, &u.Lastname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(u), nil
}

func (o *UserRepo) Create(ctx context.Context, item *domain.User) error {
	const query = `
        INSERT INTO users (telegram_id, username, firstname, lastname)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	return o.db.
		QueryRowContext(ctx, query, item.TelegramID, item.Username, item.Firstname, item.Lastname).
		Scan(&item.ID)
}

func (o *UserRepo) Update(ctx context.Context, item *domain.User) error {
	const query = `
        UPDATE users
        SET telegram_id = $1,
            username = $2,
            firstname = $3,
            lastname = $4
        WHERE id = $5
    `
	res, err := o.db.ExecContext(ctx, query, item.TelegramID, item.Username, item.Firstname, item.Lastname, item.ID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

func (o *UserRepo) Delete(ctx context.Context, id int64) error {
	const query = `
        DELETE FROM users
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

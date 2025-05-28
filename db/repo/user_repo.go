package repo

import (
	"context"
	"database/sql"
	"errors"

	"tg-star-shop-bot-001/common/domain"
)

type UserRepo struct {
	db                  *sql.DB
	adapter             *userAdapter
	userSettingsAdapter *userSettingsAdapter
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db:                  db,
		adapter:             newUserAdapter(),
		userSettingsAdapter: newUserSettingsAdapter(),
	}
}

func (o *UserRepo) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	const query = `
        SELECT *
        FROM users
        WHERE id = $1
    `
	row := o.db.QueryRowContext(ctx, query, id)
	item := &user{}
	if err := row.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt, &item.TelegramID, &item.Username, &item.FirstName, &item.LastName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(item), nil
}

func (o *UserRepo) FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	const query = `
        SELECT *
        FROM users
        WHERE telegram_id = $1
    `
	row := o.db.QueryRowContext(ctx, query, telegramID)
	item := &user{}
	if err := row.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt, &item.TelegramID, &item.Username, &item.FirstName, &item.LastName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(item), nil
}

func (o *UserRepo) Create(ctx context.Context, item *domain.User) error {
	const query = `
        INSERT INTO users (telegram_id, username, first_name, last_name)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	return o.db.
		QueryRowContext(ctx, query, item.TelegramID, item.Username, item.FirstName, item.LastName).
		Scan(&item.ID)
}

func (o *UserRepo) Update(ctx context.Context, item *domain.User) error {
	const query = `
        UPDATE users
        SET telegram_id = $1,
            username = $2,
            first_name = $3,
            last_name = $4
        WHERE id = $5
    `
	res, err := o.db.ExecContext(ctx, query, item.TelegramID, item.Username, item.FirstName, item.LastName, item.ID)
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

func (o *UserRepo) CreateSettings(ctx context.Context, userID int64) error {
	const query = `
        INSERT INTO user_settings (user_id)
        VALUES ($1)
    `
	_, err := o.db.ExecContext(ctx, query, userID)
	return err
}

func (o *UserRepo) GetSettings(ctx context.Context, userID int64) (*domain.UserSettings, error) {
	const query = `
        SELECT us.*, r.code AS role_code
        FROM user_settings us
        JOIN roles r ON us.role_id = r.id
        WHERE us.user_id = $1
    `
	row := o.db.QueryRowContext(ctx, query, userID)
	item := &userSettings{}
	if err := row.Scan(&item.UserID, &item.CreatedAt, &item.UpdatedAt, &item.Lang, &item.RoleID, &item.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.userSettingsAdapter.ToDomain(item), nil
}

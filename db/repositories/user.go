package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/meesooqa/storeque/common/domain"
	"github.com/meesooqa/storeque/db/entities"
)

// UserRepository implements the domain.UserRepository interface
type UserRepository struct {
	db                  *sql.DB
	adapter             *entities.UserAdapter
	userSettingsAdapter *entities.UserSettingsAdapter
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:                  db,
		adapter:             entities.NewUserAdapter(),
		userSettingsAdapter: entities.NewUserSettingsAdapter(),
	}
}

// FindByChatID retrieves a user by their chat ID
func (o *UserRepository) FindByChatID(ctx context.Context, chatID int64) (*domain.User, error) {
	const query = `
        SELECT *
        FROM users
        WHERE chat_id = $1
    `
	row := o.db.QueryRowContext(ctx, query, chatID)
	item := &entities.User{}
	if err := row.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt, &item.TelegramID, &item.ChatID, &item.Username, &item.FirstName, &item.LastName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(item), nil
}

// Create inserts a new user into the database
func (o *UserRepository) Create(ctx context.Context, item *domain.User) error {
	const query = `
        INSERT INTO users (chat_id, username, first_name, last_name)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	return o.db.
		QueryRowContext(ctx, query, item.ChatID, item.Username, item.FirstName, item.LastName).
		Scan(&item.ID)
}

// Update updates an existing user in the database
func (o *UserRepository) Update(ctx context.Context, item *domain.User) error {
	const query = `
        UPDATE users
        SET chat_id = $1,
            username = $2,
            first_name = $3,
            last_name = $4
        WHERE id = $5
    `
	res, err := o.db.ExecContext(ctx, query, item.ChatID, item.Username, item.FirstName, item.LastName, item.ID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

// Delete removes a user from the database by their ID
func (o *UserRepository) Delete(ctx context.Context, id int64) error {
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

// CreateSettings creates a new user settings entry in the database
func (o *UserRepository) CreateSettings(ctx context.Context, userID int64) error {
	const query = `
        INSERT INTO user_settings (user_id)
        VALUES ($1)
    `
	_, err := o.db.ExecContext(ctx, query, userID)
	return err
}

// GetSettings retrieves user settings by user ID
func (o *UserRepository) GetSettings(ctx context.Context, userID int64) (*domain.UserSettings, error) {
	const query = `
        SELECT us.*, r.code AS role_code
        FROM user_settings us
        JOIN roles r ON us.role_id = r.id
        WHERE us.user_id = $1
    `
	row := o.db.QueryRowContext(ctx, query, userID)
	item := &entities.UserSettings{}
	if err := row.Scan(&item.UserID, &item.CreatedAt, &item.UpdatedAt, &item.Lang, &item.RoleID, &item.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.userSettingsAdapter.ToDomain(item), nil
}

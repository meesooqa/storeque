package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/meesooqa/storeque/common/domain"
	"github.com/meesooqa/storeque/db/entities"
)

// UserSettingsRepository implements the domain.UserSettingsRepository interface
type UserSettingsRepository struct {
	db      *sql.DB
	adapter *entities.UserSettingsAdapter
}

// NewUserSettingsRepository creates a new instance of UserSettingsRepository
func NewUserSettingsRepository(db *sql.DB) *UserSettingsRepository {
	return &UserSettingsRepository{
		db:      db,
		adapter: entities.NewUserSettingsAdapter(),
	}
}

// FindByChatID retrieves user settings by chat ID
func (o *UserSettingsRepository) FindByChatID(ctx context.Context, chatID int64) (*domain.UserSettings, error) {
	const query = `
        SELECT us.*, r.code AS role_code
        FROM user_settings us
        JOIN users u ON u.id = us.user_id
        JOIN roles r ON r.id = us.role_id
        WHERE u.chat_id = $1
    `
	row := o.db.QueryRowContext(ctx, query, chatID)
	item := &entities.UserSettings{}
	if err := row.Scan(&item.UserID, &item.CreatedAt, &item.UpdatedAt, &item.Lang, &item.RoleID, &item.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return o.adapter.ToDomain(item), nil
}

// UpdateLangByChatID updates the language setting for a user by their chat ID
func (o *UserSettingsRepository) UpdateLangByChatID(ctx context.Context, chatID int64, value string) error {
	const query = `
        UPDATE user_settings
        SET lang = $1
        WHERE user_id = (
            SELECT id FROM users WHERE chat_id = $2
        )
    `
	res, err := o.db.ExecContext(ctx, query, value, chatID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

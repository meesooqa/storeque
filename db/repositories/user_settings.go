package repositories

import (
	"context"
	"database/sql"
	"errors"

	"tg-star-shop-bot-001/common/domain"
	"tg-star-shop-bot-001/db/entities"
)

type UserSettingsRepository struct {
	db      *sql.DB
	adapter *entities.UserSettingsAdapter
}

func NewUserSettingsRepository(db *sql.DB) *UserSettingsRepository {
	return &UserSettingsRepository{
		db:      db,
		adapter: entities.NewUserSettingsAdapter(),
	}
}

func (this *UserSettingsRepository) FindByChatID(ctx context.Context, chatID int64) (*domain.UserSettings, error) {
	const query = `
        SELECT us.*, r.code AS role_code
        FROM user_settings us
        JOIN users u ON u.id = us.user_id
        JOIN roles r ON r.id = us.role_id
        WHERE u.chat_id = $1
    `
	row := this.db.QueryRowContext(ctx, query, chatID)
	item := &entities.UserSettings{}
	if err := row.Scan(&item.UserID, &item.CreatedAt, &item.UpdatedAt, &item.Lang, &item.RoleID, &item.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // ErrNotFound
		}
		return nil, err
	}
	return this.adapter.ToDomain(item), nil
}

func (this *UserSettingsRepository) UpdateLangByChatID(ctx context.Context, chatID int64, value string) error {
	const query = `
        UPDATE user_settings
        SET lang = $1
        WHERE user_id = (
            SELECT id FROM users WHERE chat_id = $2
        )
    `
	res, err := this.db.ExecContext(ctx, query, value, chatID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows // ErrNotFound
	}
	return nil
}

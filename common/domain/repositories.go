package domain

import "context"

// UserRepository defines methods for managing user data in the database.
type UserRepository interface {
	Create(ctx context.Context, item *User) error
	Update(ctx context.Context, item *User) error
	Delete(ctx context.Context, id int64) error

	FindByChatID(ctx context.Context, chatID int64) (*User, error)

	CreateSettings(ctx context.Context, userID int64) error
	GetSettings(ctx context.Context, userID int64) (*UserSettings, error)
}

// UserSettingsRepository defines methods for managing user settings in the database.
type UserSettingsRepository interface {
	FindByChatID(ctx context.Context, chatID int64) (*UserSettings, error)
	UpdateLangByChatID(ctx context.Context, chatID int64, value string) error
}

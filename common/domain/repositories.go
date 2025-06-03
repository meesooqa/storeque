package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, item *User) error
	Update(ctx context.Context, item *User) error
	Delete(ctx context.Context, id int64) error

	FindByChatID(ctx context.Context, chatID int64) (*User, error)

	CreateSettings(ctx context.Context, userID int64) error
	GetSettings(ctx context.Context, userID int64) (*UserSettings, error)
}

type UserSettingsRepository interface {
	FindByChatID(ctx context.Context, chatID int64) (*UserSettings, error)
	UpdateLangByChatID(ctx context.Context, chatID int64, value string) error
}

type CommandRepository interface {
	FindByRoleID(ctx context.Context, roleID int64) ([]*Command, error)
}

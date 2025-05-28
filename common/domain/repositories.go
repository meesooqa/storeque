package domain

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, item *User) error
	Update(ctx context.Context, item *User) error
	Delete(ctx context.Context, id int64) error

	FindByTelegramID(ctx context.Context, telegramID int64) (*User, error)

	CreateSettings(ctx context.Context, userID int64) error
	GetSettings(ctx context.Context, userID int64) (*UserSettings, error)
}

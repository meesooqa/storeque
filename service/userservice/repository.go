package userservice

import (
	"context"

	"tg-star-shop-bot-001/common/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*domain.User, error)
	Create(ctx context.Context, item *domain.User) error
	Update(ctx context.Context, item *domain.User) error
	Delete(ctx context.Context, id int64) error

	FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error)

	CreateSettings(ctx context.Context, userID int64) error
	GetSettings(ctx context.Context, userID int64) (*domain.UserSettings, error)
}

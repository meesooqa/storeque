package userservice

import (
	"context"

	"github.com/meesooqa/storeque/common/domain"
)

// UserService defines the interface for user-related operations
type UserService interface {
	Register(ctx context.Context, item *domain.User) error
	GetUserSettings(ctx context.Context, chatID int64) (*domain.UserSettings, error)
	SetChatLang(ctx context.Context, chatID int64, value string) error
}

package userservice

import (
	"context"

	"github.com/meesooqa/storeque/common/domain"
)

// Service implements the UserService interface
type Service struct {
	userRepo         domain.UserRepository
	userSettingsRepo domain.UserSettingsRepository
}

// NewService creates a new instance of Service with the provided repositories
func NewService(userRepo domain.UserRepository, userSettingsRepo domain.UserSettingsRepository) *Service {
	return &Service{
		userRepo:         userRepo,
		userSettingsRepo: userSettingsRepo,
	}
}

// Register registers a new user in the system
func (o Service) Register(ctx context.Context, item *domain.User) error {
	existing, err := o.userRepo.FindByChatID(ctx, item.ChatID)
	if err == nil && existing != nil {
		return nil
	}

	err = o.userRepo.Create(ctx, item)
	if err != nil {
		return err
	}

	err = o.userRepo.CreateSettings(ctx, item.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserSettings retrieves the user settings for a given chat ID
func (o Service) GetUserSettings(ctx context.Context, chatID int64) (*domain.UserSettings, error) {
	return o.userSettingsRepo.FindByChatID(ctx, chatID)
}

// SetChatLang updates the language setting for a user identified by chat ID
func (o Service) SetChatLang(ctx context.Context, chatID int64, value string) error {
	return o.userSettingsRepo.UpdateLangByChatID(ctx, chatID, value)
}

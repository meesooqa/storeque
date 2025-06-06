package userservice

import (
	"context"

	"github.com/meesooqa/storeque/common/domain"
)

type Service struct {
	userRepo         domain.UserRepository
	userSettingsRepo domain.UserSettingsRepository
}

func NewService(userRepo domain.UserRepository, userSettingsRepo domain.UserSettingsRepository) *Service {
	return &Service{
		userRepo:         userRepo,
		userSettingsRepo: userSettingsRepo,
	}
}

func (this Service) Register(ctx context.Context, item *domain.User) error {
	existing, err := this.userRepo.FindByChatID(ctx, item.ChatID)
	if err == nil && existing != nil {
		return nil
	}

	err = this.userRepo.Create(ctx, item)
	if err != nil {
		return err
	}

	err = this.userRepo.CreateSettings(ctx, item.ID)
	if err != nil {
		return err
	}

	return nil
}

func (this Service) GetUserSettings(ctx context.Context, chatID int64) (*domain.UserSettings, error) {
	return this.userSettingsRepo.FindByChatID(ctx, chatID)
}

func (this Service) SetChatLang(ctx context.Context, chatID int64, value string) error {
	return this.userSettingsRepo.UpdateLangByChatID(ctx, chatID, value)
}

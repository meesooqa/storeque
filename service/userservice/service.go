package userservice

import (
	"context"

	"github.com/meesooqa/storeque/common/domain"
)

type Service struct {
	userRepo         domain.UserRepository
	userSettingsRepo domain.UserSettingsRepository
	commandRepo      domain.CommandRepository
}

func NewService(userRepo domain.UserRepository, userSettingsRepo domain.UserSettingsRepository, commandRepo domain.CommandRepository) *Service {
	return &Service{
		userRepo:         userRepo,
		userSettingsRepo: userSettingsRepo,
		commandRepo:      commandRepo,
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

func (this Service) GetUserAllowedCommands(ctx context.Context, chatID int64) ([]string, error) {
	userSettings, err := this.userSettingsRepo.FindByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	roleCommands, err := this.commandRepo.FindByRoleID(ctx, userSettings.RoleID)
	if err != nil {
		return nil, err
	}
	allowedCommands := make([]string, len(roleCommands))
	for _, roleCommand := range roleCommands {
		allowedCommands = append(allowedCommands, roleCommand.Code)
	}
	return allowedCommands, nil
}

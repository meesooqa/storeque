package userservice

import (
	"context"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/common/domain"
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

func (this *Service) Register(ctx context.Context, item *domain.User) error {
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

func (this *Service) SetChatLang(ctx context.Context, chatID int64, value string) error {
	appDeps := app.GetInstance()
	appDeps.ChangeLang(value)
	return this.userSettingsRepo.UpdateLangByChatID(ctx, chatID, value)
}

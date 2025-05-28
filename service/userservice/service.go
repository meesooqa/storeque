package userservice

import (
	"context"

	"tg-star-shop-bot-001/common/domain"
)

type Service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) *Service {
	return &Service{userRepo: userRepo}
}

func (o *Service) Register(ctx context.Context, item *domain.User) error {
	existing, err := o.userRepo.FindByTelegramID(ctx, item.TelegramID)
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

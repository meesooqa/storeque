package userservice

import (
	"context"

	"tg-star-shop-bot-001/common/domain"
)

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (o *Service) Register(ctx context.Context, item *domain.User) error {
	existing, err := o.repository.FindByTelegramID(ctx, item.TelegramID)
	if err == nil && existing != nil {
		return nil
	}

	err = o.repository.Create(ctx, item)
	if err != nil {
		return err
	}

	err = o.repository.CreateSettings(ctx, item.ID)
	if err != nil {
		return err
	}

	return nil
}

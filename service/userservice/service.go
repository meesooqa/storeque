package userservice

import (
	"context"

	"tg-star-shop-bot-001/common/domain"
	"tg-star-shop-bot-001/db/repo"
)

type Service struct {
	repo *repo.UserRepo
}

func NewService(repo *repo.UserRepo) *Service {
	return &Service{repo: repo}
}

func (o *Service) Register(ctx context.Context, item *domain.User) error {
	err := o.repo.Create(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

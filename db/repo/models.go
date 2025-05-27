package repo

import (
	"time"

	"tg-star-shop-bot-001/common/domain"
)

type user struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	TelegramID int64     `db:"telegram_id"`
	Username   string    `db:"username"`
	Firstname  string    `db:"firstname"`
	Lastname   string    `db:"lastname"`
}

type userAdapter struct{}

func newUserAdapter() *userAdapter {
	return &userAdapter{}
}

func (a *userAdapter) ToDomain(item *user) *domain.User {
	return &domain.User{
		ID:         item.ID,
		Username:   item.Username,
		Firstname:  item.Firstname,
		Lastname:   item.Lastname,
		TelegramID: item.TelegramID,
	}
}

func (a *userAdapter) FromDomain(item *domain.User) *user {
	return &user{
		ID:         item.ID,
		Username:   item.Username,
		Firstname:  item.Firstname,
		Lastname:   item.Lastname,
		TelegramID: item.TelegramID,
	}
}

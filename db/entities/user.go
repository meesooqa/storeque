package entities

import (
	"time"

	"tg-star-shop-bot-001/common/domain"
)

type User struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	TelegramID int64     `db:"telegram_id"`
	ChatID     int64     `db:"chat_id"`
	Username   string    `db:"username"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
}

type UserAdapter struct{}

func NewUserAdapter() *UserAdapter {
	return &UserAdapter{}
}

func (a *UserAdapter) ToDomain(item *User) *domain.User {
	return &domain.User{
		ID:         item.ID,
		TelegramID: item.TelegramID,
		ChatID:     item.ChatID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
	}
}

func (a *UserAdapter) FromDomain(item *domain.User) *User {
	return &User{
		ID:         item.ID,
		TelegramID: item.TelegramID,
		ChatID:     item.ChatID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
	}
}

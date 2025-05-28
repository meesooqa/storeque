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
	Username   string    `db:"username"`
	FirstName  string    `db:"firstname"`
	LastName   string    `db:"lastname"`
}

type UserAdapter struct{}

func NewUserAdapter() *UserAdapter {
	return &UserAdapter{}
}

func (a *UserAdapter) ToDomain(item *User) *domain.User {
	return &domain.User{
		ID:         item.ID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
		TelegramID: item.TelegramID,
	}
}

func (a *UserAdapter) FromDomain(item *domain.User) *User {
	return &User{
		ID:         item.ID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
		TelegramID: item.TelegramID,
	}
}

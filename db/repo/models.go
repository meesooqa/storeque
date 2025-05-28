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
	FirstName  string    `db:"firstname"`
	LastName   string    `db:"lastname"`
}

type userAdapter struct{}

func newUserAdapter() *userAdapter {
	return &userAdapter{}
}

func (a *userAdapter) ToDomain(item *user) *domain.User {
	return &domain.User{
		ID:         item.ID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
		TelegramID: item.TelegramID,
	}
}

func (a *userAdapter) FromDomain(item *domain.User) *user {
	return &user{
		ID:         item.ID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
		TelegramID: item.TelegramID,
	}
}

type role struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Code      string    `db:"code"`
}

type roleAdapter struct{}

func newRoleAdapter() *roleAdapter {
	return &roleAdapter{}
}

func (a *roleAdapter) ToDomain(item *role) *domain.Role {
	return &domain.Role{
		ID:   item.ID,
		Code: item.Code,
	}
}

func (a *roleAdapter) FromDomain(item *domain.Role) *role {
	return &role{
		ID:   item.ID,
		Code: item.Code,
	}
}

package entities

import (
	"time"

	"github.com/meesooqa/storeque/common/domain"
)

// User represents a user in the database
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

// UserAdapter converts between User domain model and User database entity
type UserAdapter struct{}

// NewUserAdapter creates a new instance of UserAdapter
func NewUserAdapter() *UserAdapter {
	return &UserAdapter{}
}

// ToDomain converts a User database entity to a User domain model
func (o *UserAdapter) ToDomain(item *User) *domain.User {
	return &domain.User{
		ID:         item.ID,
		TelegramID: item.TelegramID,
		ChatID:     item.ChatID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
	}
}

// FromDomain converts a User domain model to a User database entity
func (o *UserAdapter) FromDomain(item *domain.User) *User {
	return &User{
		ID:         item.ID,
		TelegramID: item.TelegramID,
		ChatID:     item.ChatID,
		Username:   item.Username,
		FirstName:  item.FirstName,
		LastName:   item.LastName,
	}
}

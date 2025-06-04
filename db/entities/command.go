package entities

import (
	"time"

	"github.com/meesooqa/tg-star-shop-bot/common/domain"
)

type Command struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Code      string    `db:"code"`
	RoleID    int64     `db:"role_id"`
	Role      string    `db:"role_code"` // SELECT us.*, r.code AS role_code
}

type CommandAdapter struct{}

func NewCommandAdapter() *CommandAdapter {
	return &CommandAdapter{}
}

func (this CommandAdapter) ToDomain(item *Command) *domain.Command {
	return &domain.Command{
		ID:     item.ID,
		Code:   item.Code,
		RoleID: item.RoleID,
		Role: &domain.Role{
			ID:   item.RoleID,
			Code: item.Role,
		},
	}
}

func (this CommandAdapter) FromDomain(item *domain.Command) *Command {
	return &Command{
		ID:     item.ID,
		Code:   item.Code,
		RoleID: item.RoleID,
		Role:   item.Role.Code,
	}
}

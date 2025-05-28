package entities

import (
	"time"

	"tg-star-shop-bot-001/common/domain"
)

type Role struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Code      string    `db:"code"`
}

type RoleAdapter struct{}

func NewRoleAdapter() *RoleAdapter {
	return &RoleAdapter{}
}

func (a *RoleAdapter) ToDomain(item *Role) *domain.Role {
	return &domain.Role{
		ID:   item.ID,
		Code: item.Code,
	}
}

func (a *RoleAdapter) FromDomain(item *domain.Role) *Role {
	return &Role{
		ID:   item.ID,
		Code: item.Code,
	}
}

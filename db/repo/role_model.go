package repo

import (
	"time"

	"tg-star-shop-bot-001/common/domain"
)

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

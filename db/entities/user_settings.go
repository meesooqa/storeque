package entities

import (
	"time"

	"github.com/meesooqa/tg-star-shop-bot/common/domain"
)

type UserSettings struct {
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	RoleID    int64     `db:"role_id"`
	Lang      string    `db:"lang"`
	Role      string    `db:"role_code"` // SELECT us.*, r.code AS role_code
}

type UserSettingsAdapter struct{}

func NewUserSettingsAdapter() *UserSettingsAdapter {
	return &UserSettingsAdapter{}
}

func (a *UserSettingsAdapter) ToDomain(item *UserSettings) *domain.UserSettings {
	return &domain.UserSettings{
		UserID: item.UserID,
		RoleID: item.RoleID,
		Lang:   item.Lang,
		Role: &domain.Role{
			ID:   item.RoleID,
			Code: item.Role,
		},
	}
}

func (a *UserSettingsAdapter) FromDomain(item *domain.UserSettings) *UserSettings {
	return &UserSettings{
		UserID: item.UserID,
		RoleID: item.RoleID,
		Role:   item.Role.Code,
		Lang:   item.Lang,
	}
}

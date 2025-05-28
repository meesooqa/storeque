package repo

import (
	"time"

	"tg-star-shop-bot-001/common/domain"
)

type userSettings struct {
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	RoleID    int64     `db:"role_id"`
	Lang      string    `db:"lang"`
	Role      string    `db:"role_code"` // SELECT us.*, r.code AS role_code
}

type userSettingsAdapter struct{}

func newUserSettingsAdapter() *userSettingsAdapter {
	return &userSettingsAdapter{}
}

func (a *userSettingsAdapter) ToDomain(item *userSettings) *domain.UserSettings {
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

func (a *userSettingsAdapter) FromDomain(item *domain.UserSettings) *userSettings {
	return &userSettings{
		UserID: item.UserID,
		RoleID: item.RoleID,
		Role:   item.Role.Code,
		Lang:   item.Lang,
	}
}

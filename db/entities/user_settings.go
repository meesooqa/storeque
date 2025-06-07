package entities

import (
	"time"

	"github.com/meesooqa/storeque/common/domain"
)

// UserSettings represents user settings in the database
type UserSettings struct {
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	RoleID    int64     `db:"role_id"`
	Lang      string    `db:"lang"`
	Role      string    `db:"role_code"` // SELECT us.*, r.code AS role_code
}

// UserSettingsAdapter converts between UserSettings domain model and UserSettings database entity
type UserSettingsAdapter struct{}

// NewUserSettingsAdapter creates a new instance of UserSettingsAdapter
func NewUserSettingsAdapter() *UserSettingsAdapter {
	return &UserSettingsAdapter{}
}

// ToDomain converts a UserSettings database entity to a UserSettings domain model
func (o *UserSettingsAdapter) ToDomain(item *UserSettings) *domain.UserSettings {
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

// FromDomain converts a UserSettings domain model to a UserSettings database entity
func (o *UserSettingsAdapter) FromDomain(item *domain.UserSettings) *UserSettings {
	return &UserSettings{
		UserID: item.UserID,
		RoleID: item.RoleID,
		Role:   item.Role.Code,
		Lang:   item.Lang,
	}
}

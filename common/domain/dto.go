package domain

const (
	// UserSettingsLangValueEn is English language code
	UserSettingsLangValueEn = "en"

	// UserSettingsLangValueRu is Russian language code
	UserSettingsLangValueRu = "ru"
)

// User represents a user in the system with their Telegram ID, chat ID, and personal details.
type User struct {
	ID         int64
	TelegramID int64
	ChatID     int64
	Username   string
	FirstName  string
	LastName   string
	Settings   *UserSettings
}

// Role represents a user role in the system with an ID and a code.
type Role struct {
	ID   int64
	Code string
}

// UserSettings represents the settings associated with a user, including their role and language preference.
type UserSettings struct {
	UserID int64
	RoleID int64
	Role   *Role
	Lang   string
}

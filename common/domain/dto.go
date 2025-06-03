package domain

const (
	UserSettingsLangValueEn = "en"
	UserSettingsLangValueRu = "ru"
)

type User struct {
	ID         int64
	TelegramID int64
	ChatID     int64
	Username   string
	FirstName  string
	LastName   string
	Settings   *UserSettings
}

type Role struct {
	ID   int64
	Code string
}

type UserSettings struct {
	UserID int64
	RoleID int64
	Role   *Role
	Lang   string
}

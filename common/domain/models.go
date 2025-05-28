package domain

type User struct {
	ID         int64
	TelegramID int64
	Username   string
	FirstName  string
	LastName   string
	Settings   UserSettings
}

type Role struct {
	ID   int64
	Code string
}

type UserSettings struct {
	Role string
	Lang string
}

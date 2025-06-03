package entities

type RoleCommand struct {
	RoleID    int64 `db:"role_id"`
	CommandID int64 `db:"command_id"`
}

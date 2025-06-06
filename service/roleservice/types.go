package roleservice

import "context"

type RoleService interface {
	GetAllowedCommands(ctx context.Context, role string) []string
	IsCommandAllowed(ctx context.Context, role, command string) bool
}

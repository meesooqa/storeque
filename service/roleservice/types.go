package roleservice

import "context"

// RoleService defines the interface for role-based command access control
type RoleService interface {
	GetAllowedCommands(ctx context.Context, role string) []string
	IsCommandAllowed(ctx context.Context, role, command string) bool
}

package roleservice

import "context"

// Service implements the RoleService interface
type Service struct {
	roleCommands map[string][]string
}

// NewService creates a new instance of Service with predefined role commands
func NewService() *Service {
	return &Service{
		roleCommands: map[string][]string{
			"admin":    {"start", "help", "settings", "buy", "dice"},
			"customer": {"start", "help", "buy", "dice"},
		},
	}
}

// IsCommandAllowed checks if a command is allowed for a given role
func (o Service) IsCommandAllowed(ctx context.Context, role, command string) bool {
	allowedCommands := o.GetAllowedCommands(ctx, role)
	if allowedCommands == nil {
		return false
	}
	for _, allowedCommand := range allowedCommands {
		if allowedCommand == command {
			return true
		}
	}
	return false
}

// GetAllowedCommands retrieves the list of commands allowed for a given role
func (o Service) GetAllowedCommands(_ context.Context, role string) []string {
	allowedCommands, exists := o.roleCommands[role]
	if !exists {
		return nil
	}
	return allowedCommands
}

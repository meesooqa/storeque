package roleservice

import "context"

type Service struct {
	roleCommands map[string][]string
}

func NewService() *Service {
	return &Service{
		roleCommands: map[string][]string{
			"admin":    {"start", "help", "settings", "buy", "dice"},
			"customer": {"start", "help", "buy", "dice"},
		},
	}
}

func (this Service) IsCommandAllowed(ctx context.Context, role string, command string) bool {
	allowedCommands := this.GetAllowedCommands(ctx, role)
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

func (this Service) GetAllowedCommands(ctx context.Context, role string) []string {
	allowedCommands, exists := this.roleCommands[role]
	if !exists {
		//return []string{}
		return nil
	}
	return allowedCommands
}

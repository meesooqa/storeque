package middleware

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/domain"
	"github.com/meesooqa/storeque/service/userservice"
)

// Register is a middleware that registers a user when they send a message
type Register struct {
	userService userservice.UserService
}

// NewRegister creates a new Register middleware
func NewRegister(userService userservice.UserService) *Register {
	return &Register{
		userService: userService,
	}
}

// Execute registers the user based on the update received
func (o *Register) Execute(ctx context.Context, update *tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}
	// TODO cache
	return o.userService.Register(ctx, &domain.User{
		ChatID: update.Message.From.ID,
		// TelegramID: update.Message.From.ID,
		Username:  update.Message.From.UserName,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	})
}

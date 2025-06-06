package middleware

import (
	"context"
	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/meesooqa/storeque/service/userservice"

	"github.com/meesooqa/storeque/common/domain"
)

type Register struct {
	userService userservice.UserService
}

func NewRegister(userService userservice.UserService) *Register {
	return &Register{
		userService: userService,
	}
}

// TODO cache
func (this *Register) Execute(ctx context.Context, update *tgbotapi.Update) error {
	if update.Message == nil {
		//return fmt.Errorf("message is nil")
		return nil
	}
	return this.userService.Register(ctx, &domain.User{
		ChatID: update.Message.From.ID,
		// TelegramID: update.Message.From.ID,
		Username:  update.Message.From.UserName,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	})
}

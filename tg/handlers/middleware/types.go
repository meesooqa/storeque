package middleware

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

type UpdatePreHandler interface {
	Execute(ctx context.Context, update *tgbotapi.Update) error
}

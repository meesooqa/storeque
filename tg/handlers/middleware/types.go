package middleware

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

// UpdatePreHandler is an interface for middleware that processes updates before they are handled by the main logic
type UpdatePreHandler interface {
	Execute(ctx context.Context, update *tgbotapi.Update) error
}

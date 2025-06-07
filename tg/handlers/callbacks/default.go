package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

// DefaultHandler is a callback handler that serves as a fallback for unhandled callbacks
type DefaultHandler struct {
	BaseHandler
}

// NewDefaultHandler creates a new DefaultHandler with the provided application dependencies and bot API
func NewDefaultHandler(appDeps app.App, bot *tgbotapi.BotAPI) *DefaultHandler {
	return &DefaultHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

// GetData returns an empty string, indicating that this handler does not have a specific data identifier
func (o *DefaultHandler) GetData() string {
	return ""
}

// Handle processes the callback query and sends a default response
func (o *DefaultHandler) Handle(_ context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, "Default Callback Handler")) // nolint
}

package commands

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

// DefaultHandler is a command handler that responds to unknown commands
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

// GetName returns an empty string, indicating that this handler does not have a specific command name
func (o *DefaultHandler) GetName() string {
	return ""
}

// GetDescription returns an empty string, indicating that this handler does not have a specific description
func (o *DefaultHandler) GetDescription(_ lang.Localization) string {
	return ""
}

// Handle processes the incoming message and sends a default response for unknown commands
func (o *DefaultHandler) Handle(_ context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	text := "Неизвестная команда. Используйте /help для получения списка команд."
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ReplyParameters.MessageID = inputMessage.MessageID
	o.bot.Send(msg) // nolint
}

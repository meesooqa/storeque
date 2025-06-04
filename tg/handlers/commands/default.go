package commands

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

type DefaultHandler struct {
	BaseHandler
}

func NewDefaultHandler(appDeps app.App, bot *tgbotapi.BotAPI) *DefaultHandler {
	return &DefaultHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o DefaultHandler) GetName() string {
	return ""
}

func (o DefaultHandler) GetDescription(loc lang.Localization) string {
	return ""
}

func (o *DefaultHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	text := "Неизвестная команда. Используйте /help для получения списка команд."
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ReplyParameters.MessageID = inputMessage.MessageID
	o.bot.Send(msg)
}

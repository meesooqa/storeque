package commands

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
)

type DefaultHandler struct {
	BaseHandler
}

func NewDefaultHandler(appDeps *app.AppDeps, bot *tgbotapi.BotAPI) *DefaultHandler {
	return &DefaultHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o *DefaultHandler) GetName() string {
	return ""
}

func (o *DefaultHandler) GetDescription() string {
	return ""
}

func (o *DefaultHandler) Handle(ctx context.Context, inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	text := "Неизвестная команда. Используйте /help для получения списка команд."
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ReplyParameters.MessageID = inputMessage.MessageID
	o.bot.Send(msg)
}

package commands

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
)

type StartHandler struct {
	BaseHandler
}

func NewStartHandler(bot *tgbotapi.BotAPI, appDeps *app.AppDeps) *StartHandler {
	return &StartHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o *StartHandler) GetName() string {
	return "start"
}

func (o *StartHandler) GetDescription() string {
	return o.appDeps.Lang.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *StartHandler) Handle(inputMessage *tgbotapi.Message) {
	// TODO save user to DB
	// TODO langTag from user config
	//  "Welcome, {{.UserName}}!"
	//  "Use /help to see commands list."
	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	//o.bot.Send(msg)
}

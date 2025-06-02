package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
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

func (o *DefaultHandler) GetData() string {
	return ""
}

func (o *DefaultHandler) GetText() string {
	return ""
}

func (o *DefaultHandler) Render() {
}

func (o *DefaultHandler) Handle(ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) {
	o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, "Default Callback Handler"))
}

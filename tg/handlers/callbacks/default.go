package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/tg-star-shop-bot/common/app"
	"github.com/meesooqa/tg-star-shop-bot/common/lang"
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

func (o *DefaultHandler) Handle(ctx context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, "Default Callback Handler"))
}

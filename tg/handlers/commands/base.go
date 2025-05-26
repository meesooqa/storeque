package commands

import (
	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
)

type CommandHandler interface {
	Handle(*tgbotapi.Message)
	GetName() string
	GetDescription() string
}

type BaseHandler struct {
	bot     *tgbotapi.BotAPI
	appDeps *app.AppDeps
}

package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/service/userservice"
)

type CallbackHandler interface {
	Handle(context.Context, *tgbotapi.CallbackQuery)
	GetData() string
}

type BaseHandler struct {
	bot      *tgbotapi.BotAPI
	appDeps  app.App
	children []CallbackHandler
}

// GetAll returns list of all callbacks
func GetAll(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) map[string]CallbackHandler {
	list := []CallbackHandler{
		NewLangRuHandler(appDeps, bot, userService),
		NewLangEnHandler(appDeps, bot, userService),
	}

	handlersMap := make(map[string]CallbackHandler)
	for _, item := range list {
		if item.GetData() != "" {
			handlersMap[item.GetData()] = item
		}
	}
	return handlersMap
}

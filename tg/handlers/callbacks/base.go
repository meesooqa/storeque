package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

type CallbackHandler interface {
	Handle(context.Context, lang.Localization, *tgbotapi.CallbackQuery)
	GetData() string
}

type BaseHandler struct {
	bot     *tgbotapi.BotAPI
	appDeps app.App
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

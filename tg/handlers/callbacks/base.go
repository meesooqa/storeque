package callbacks

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// CallbackHandler is an interface for handling callback queries in Telegram bot
type CallbackHandler interface {
	Handle(context.Context, lang.Localization, *tgbotapi.CallbackQuery)
	GetData() string
}

// BaseHandler is a base struct for callback handlers, providing common dependencies
type BaseHandler struct {
	appDeps app.App
	bot     *tgbotapi.BotAPI
}

// GetAll returns a map of all callback handlers, keyed by their data string
func GetAll(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) map[string]CallbackHandler {
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

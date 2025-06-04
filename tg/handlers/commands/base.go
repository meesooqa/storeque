package commands

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

type CommandHandler interface {
	Handle(context.Context, lang.Localization, *tgbotapi.Message)
	GetName() string
	GetDescription(loc lang.Localization) string
}

type BaseHandler struct {
	bot     *tgbotapi.BotAPI
	loc     lang.Localization
	appDeps app.App
}

// GetAll returns list of all commands
func GetAll(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) map[string]CommandHandler {
	help := NewHelpHandler(appDeps, bot)
	list := []CommandHandler{
		NewStartHandler(appDeps, bot, userService),
		help,
		NewSettingsHandler(appDeps, bot, userService),
		NewBuyHandler(appDeps, bot),
		NewDiceHandler(appDeps, bot),
	}
	help.SetCommands(list)

	handlersMap := make(map[string]CommandHandler)
	for _, item := range list {
		if item.GetName() != "" {
			handlersMap[item.GetName()] = item
		}
	}
	return handlersMap
}

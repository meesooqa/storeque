package commands

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// CommandHandler is an interface for handling commands in Telegram bot
type CommandHandler interface {
	Handle(context.Context, lang.Localization, *tgbotapi.Message)
	GetName() string
	GetDescription(loc lang.Localization) string
}

// BaseHandler is a base struct for command handlers, providing common dependencies
type BaseHandler struct {
	appDeps app.App
	bot     *tgbotapi.BotAPI
}

// GetAll returns a map of all command handlers, keyed by their command name
func GetAll(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) map[string]CommandHandler {
	help := NewHelpHandler(appDeps, bot)
	list := []CommandHandler{
		NewStartHandler(appDeps, bot, userService),
		help,
		NewSettingsHandler(appDeps, bot, userService),
		NewBuyHandler(appDeps, bot),
		NewDiceHandler(appDeps, bot),
	}
	help.SetCommands(list) // all commands. Search "SetAllowedCommands"

	handlersMap := make(map[string]CommandHandler)
	for _, item := range list {
		if item.GetName() != "" {
			handlersMap[item.GetName()] = item
		}
	}
	return handlersMap
}

package handlers

import (
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/tg/handlers/commands"
)

func getCommandHandlers(bot *tgbotapi.BotAPI, appDeps *app.AppDeps) map[string]commands.CommandHandler {
	// TODO UserGroup
	helpHandler := commands.NewHelpHandler(bot, appDeps)
	handlers := []commands.CommandHandler{
		commands.NewStartHandler(bot, appDeps),
		helpHandler,
		commands.NewBuyHandler(bot, appDeps),
		// commands.NewMyHandler(bot, appDeps),
		commands.NewDiceHandler(bot, appDeps),
	}
	helpHandler.SetAvailableCommands(handlers)

	handlersMap := make(map[string]commands.CommandHandler)
	for _, h := range handlers {
		handlersMap[h.GetName()] = h
	}
	return handlersMap
}

type TelegramHandler struct {
	commands map[string]commands.CommandHandler
	bot      *tgbotapi.BotAPI
	appDeps  *app.AppDeps
}

func NewTelegramHandler(bot *tgbotapi.BotAPI, appDeps *app.AppDeps) *TelegramHandler {
	return &TelegramHandler{
		commands: getCommandHandlers(bot, appDeps),
		bot:      bot,
		appDeps:  appDeps,
	}
}

func (o *TelegramHandler) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			o.appDeps.Logger.Info("recovered from panic", slog.Any("panicValue", panicValue))
		}
	}()

	if update.Message == nil {
		return
	}

	o.appDeps.Logger.Debug("update.Message", slog.String("UserName", "update.Message.From.UserName"), slog.String("Text", update.Message.Text))

	// TODO SuccessfulPayment
	//if update.Message.SuccessfulPayment != nil {
	//	handleSuccessfulPayment(update.Message)
	//	return
	//}

	if command, ok := o.commands[update.Message.Command()]; ok {
		command.Handle(update.Message)
	} else {
		cmdHandler := commands.NewDefaultHandler(o.bot, o.appDeps)
		cmdHandler.Handle(update.Message)
	}
}

package handlers

import (
	"context"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/service/userservice"
	"tg-star-shop-bot-001/tg/handlers/callbacks"
	"tg-star-shop-bot-001/tg/handlers/commands"
)

func getCommandHandlers(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) map[string]commands.CommandHandler {
	// TODO filter by UserGroup
	handlersMap := commands.GetAll(appDeps, bot, userService)
	return handlersMap
}

func getCallbackHandlers(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) map[string]callbacks.CallbackHandler {
	handlersMap := callbacks.GetAll(appDeps, bot, userService)
	return handlersMap
}

type TelegramHandler struct {
	commands  map[string]commands.CommandHandler
	callbacks map[string]callbacks.CallbackHandler
	bot       *tgbotapi.BotAPI
	appDeps   app.App
}

func NewTelegramHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *TelegramHandler {
	return &TelegramHandler{
		commands:  getCommandHandlers(appDeps, bot, userService),
		callbacks: getCallbackHandlers(appDeps, bot, userService),
		bot:       bot,
		appDeps:   appDeps,
	}
}

func (o *TelegramHandler) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			o.appDeps.Logger().Info("recovered from panic", slog.Any("panicValue", panicValue))
		}
	}()

	//if !CheckAuthorizedUser(update.Message.From.ID) {
	//	return
	//}

	// TODO SuccessfulPayment
	/*if update.Message.SuccessfulPayment != nil {
		handleSuccessfulPayment(update.Message)
		return
	}*/
	if update.CallbackQuery != nil {
		o.appDeps.Logger().Info("clbk data", slog.String("data", update.CallbackQuery.Data))
		if callback, ok := o.callbacks[update.CallbackQuery.Data]; ok {
			callback.Handle(ctx, update.CallbackQuery)
		} else {
			clbkHandler := callbacks.NewDefaultHandler(o.appDeps, o.bot)
			clbkHandler.Handle(ctx, update.CallbackQuery)
		}
		return
	}

	if update.Message == nil {
		return
	}

	if command, ok := o.commands[update.Message.Command()]; ok {
		command.Handle(ctx, update.Message)
	} else {
		cmdHandler := commands.NewDefaultHandler(o.appDeps, o.bot)
		cmdHandler.Handle(ctx, update.Message)
	}
}

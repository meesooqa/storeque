package handlers

import (
	"context"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/tg/handlers/middleware"
)

// TelegramHandler is a handler for Telegram
type TelegramHandler struct {
	appDeps app.App
	bot     *tgbotapi.BotAPI

	updatePreHandlers []middleware.UpdatePreHandler

	commandRouter  *CommandRouter
	callbackRouter *CallbackRouter
}

// NewTelegramHandler creates a new TelegramHandler
func NewTelegramHandler(appDeps app.App, bot *tgbotapi.BotAPI, updateMiddleware []middleware.UpdatePreHandler, commandRouter *CommandRouter, callbackRouter *CallbackRouter) *TelegramHandler {
	return &TelegramHandler{
		appDeps:           appDeps,
		bot:               bot,
		updatePreHandlers: updateMiddleware,
		commandRouter:     commandRouter,
		callbackRouter:    callbackRouter,
	}
}

// HandleUpdate handles incoming updates from Telegram
func (o *TelegramHandler) HandleUpdate(ctx context.Context, update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			o.appDeps.Logger().Info("recovered from panic", slog.Any("panicValue", panicValue))
		}
	}()

	var err error

	for _, ph := range o.updatePreHandlers {
		if err = ph.Execute(ctx, update); err != nil {
			o.appDeps.Logger().Error("updateMiddleware", slog.Any("error", err), slog.Any("updatePreHandler", ph))
			return
		}
	}

	if err = o.callbackRouter.Route(ctx, update); err != nil {
		o.appDeps.Logger().Error("callbackRouter", slog.Any("error", err))
		return
	}

	if err = o.commandRouter.Route(ctx, update); err != nil {
		o.appDeps.Logger().Error("commandRouter", slog.Any("error", err))
		return
	}

	// TODO SuccessfulPayment
	/*if update.Message.SuccessfulPayment != nil {
		handleSuccessfulPayment(update.Message)
		return
	}*/
}

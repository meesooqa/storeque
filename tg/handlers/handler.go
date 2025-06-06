package handlers

import (
	"context"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/tg/handlers/middleware"
)

type TelegramHandler struct {
	appDeps app.App
	bot     *tgbotapi.BotAPI

	updatePreHandlers []middleware.UpdatePreHandler

	commandRouter  *CommandRouter
	callbackRouter *CallbackRouter
}

func NewTelegramHandler(appDeps app.App, bot *tgbotapi.BotAPI, updateMiddleware []middleware.UpdatePreHandler, commandRouter *CommandRouter, callbackRouter *CallbackRouter) *TelegramHandler {
	return &TelegramHandler{
		appDeps:           appDeps,
		bot:               bot,
		updatePreHandlers: updateMiddleware,
		commandRouter:     commandRouter,
		callbackRouter:    callbackRouter,
	}
}

func (this TelegramHandler) HandleUpdate(ctx context.Context, update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			this.appDeps.Logger().Info("recovered from panic", slog.Any("panicValue", panicValue))
		}
	}()

	var err error

	for _, ph := range this.updatePreHandlers {
		if err = ph.Execute(ctx, update); err != nil {
			this.appDeps.Logger().Error("updateMiddleware", slog.Any("error", err), slog.Any("updatePreHandler", ph))
			return
		}
	}

	if err = this.callbackRouter.Route(ctx, update); err != nil {
		this.appDeps.Logger().Error("callbackRouter", slog.Any("error", err))
		return
	}

	if err = this.commandRouter.Route(ctx, update); err != nil {
		this.appDeps.Logger().Error("commandRouter", slog.Any("error", err))
		return
	}

	// TODO SuccessfulPayment
	/*if update.Message.SuccessfulPayment != nil {
		handleSuccessfulPayment(update.Message)
		return
	}*/
}

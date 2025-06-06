package handlers

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/service/locservice"
	"github.com/meesooqa/storeque/service/userservice"
	"github.com/meesooqa/storeque/tg/handlers/callbacks"
)

type CallbackRouter struct {
	list           map[string]callbacks.CallbackHandler
	defaultHandler callbacks.CallbackHandler

	locService locservice.LocService
}

func NewCallbackRouter(appDeps app.App, bot *tgbotapi.BotAPI, locService locservice.LocService, userService userservice.UserService) *CallbackRouter {
	return &CallbackRouter{
		list:           callbacks.GetAll(appDeps, bot, userService),
		defaultHandler: callbacks.NewDefaultHandler(appDeps, bot),

		locService: locService,
	}
}

func (this *CallbackRouter) Route(ctx context.Context, update *tgbotapi.Update) error {
	if update.CallbackQuery == nil {
		return nil
	}

	chatID := this.chatIdFromUpdate(update)
	loc := this.locService.GetLoc(ctx, chatID)

	name := update.CallbackQuery.Data
	if name == "" {
		// not a callback
		this.defaultHandler.Handle(ctx, loc, update.CallbackQuery)
		return nil
	}

	handler, exists := this.list[name]
	if !exists {
		// callback not found
		this.defaultHandler.Handle(ctx, loc, update.CallbackQuery)
		return nil
	}

	handler.Handle(ctx, loc, update.CallbackQuery)
	return nil
}

func (this *CallbackRouter) chatIdFromUpdate(update *tgbotapi.Update) int64 {
	var chatID int64 = 0
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	}
	return chatID
}

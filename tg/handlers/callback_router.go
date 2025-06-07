package handlers

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/service/locservice"
	"github.com/meesooqa/storeque/service/userservice"
	"github.com/meesooqa/storeque/tg/handlers/callbacks"
)

// CallbackRouter routes callback queries to their respective handlers
type CallbackRouter struct {
	list           map[string]callbacks.CallbackHandler
	defaultHandler callbacks.CallbackHandler

	locService locservice.LocService
}

// NewCallbackRouter creates a new CallbackRouter with the provided dependencies
func NewCallbackRouter(appDeps app.App, bot *tgbotapi.BotAPI, locService locservice.LocService, userService userservice.UserService) *CallbackRouter {
	return &CallbackRouter{
		list:           callbacks.GetAll(appDeps, bot, userService),
		defaultHandler: callbacks.NewDefaultHandler(appDeps, bot),

		locService: locService,
	}
}

// Route routes the incoming update to the appropriate callback handler
func (o *CallbackRouter) Route(ctx context.Context, update *tgbotapi.Update) error {
	if update.CallbackQuery == nil {
		return nil
	}

	chatID := o.chatIDFromUpdate(update)
	loc := o.locService.GetLoc(ctx, chatID)

	name := update.CallbackQuery.Data
	if name == "" {
		// not a callback
		o.defaultHandler.Handle(ctx, loc, update.CallbackQuery)
		return nil
	}

	handler, exists := o.list[name]
	if !exists {
		// callback not found
		o.defaultHandler.Handle(ctx, loc, update.CallbackQuery)
		return nil
	}

	handler.Handle(ctx, loc, update.CallbackQuery)
	return nil
}

// chatIDFromUpdate extracts the chat ID from the update
func (o *CallbackRouter) chatIDFromUpdate(update *tgbotapi.Update) int64 {
	var chatID int64
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	}
	return chatID
}

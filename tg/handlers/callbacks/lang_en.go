package callbacks

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/domain"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// LangEnHandler is a callback handler for setting the language to English
type LangEnHandler struct {
	BaseHandler
	userService userservice.UserService
}

// NewLangEnHandler creates a new LangEnHandler with the provided application dependencies, bot API, and user service
func NewLangEnHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) *LangEnHandler {
	return &LangEnHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

// GetData returns the data string associated with this handler, which is used to identify it in callback queries
func (o *LangEnHandler) GetData() string {
	return fmt.Sprintf("lang-%s", domain.UserSettingsLangValueEn)
}

// Handle processes the callback query to set the user's language to English
func (o *LangEnHandler) Handle(ctx context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID

	err := o.userService.SetChatLang(ctx, chatID, domain.UserSettingsLangValueEn)
	loc.SetLang(domain.UserSettingsLangValueEn)
	if err != nil {
		o.appDeps.Logger().Error("LangEnHandler", slog.Any("error", err))
		// Remove loading animation and show popup message
		o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, loc.Localize("tg.error.updatelang", nil))) // nolint
		return
	}

	o.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.clbk.lang.en", nil))) // nolint
	// Remove loading animation
	o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, "")) // nolint
}

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

// LangRuHandler is a callback handler for setting the language to Russian
type LangRuHandler struct {
	BaseHandler
	userService userservice.UserService
}

// NewLangRuHandler creates a new LangRuHandler with the provided application dependencies, bot API, and user service
func NewLangRuHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) *LangRuHandler {
	return &LangRuHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

// GetData returns the data string associated with this handler, which is used to identify it in callback queries
func (o *LangRuHandler) GetData() string {
	return fmt.Sprintf("lang-%s", domain.UserSettingsLangValueRu)
}

// Handle processes the callback query to set the user's language to Russian
func (o *LangRuHandler) Handle(ctx context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID

	err := o.userService.SetChatLang(ctx, chatID, domain.UserSettingsLangValueRu)
	loc.SetLang(domain.UserSettingsLangValueRu)
	if err != nil {
		o.appDeps.Logger().Error("LangRuHandler", slog.Any("error", err))
		// Remove loading animation and show popup message
		o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, loc.Localize("tg.error.updatelang", nil))) // nolint
		return
	}

	o.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.clbk.lang.ru", nil))) // nolint
	// Remove loading animation
	o.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, "")) // nolint
}

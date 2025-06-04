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

type LangRuHandler struct {
	BaseHandler
	userService *userservice.Service
}

func NewLangRuHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *LangRuHandler {
	return &LangRuHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

func (this *LangRuHandler) GetData() string {
	return fmt.Sprintf("lang-%s", domain.UserSettingsLangValueRu)
}

func (this *LangRuHandler) Handle(ctx context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID

	err := this.userService.SetChatLang(ctx, chatID, domain.UserSettingsLangValueRu)
	loc.SetLang(domain.UserSettingsLangValueRu)
	if err != nil {
		this.appDeps.Logger().Error("LangRuHandler", slog.Any("error", err))
		// Remove loading animation and show popup message
		this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, loc.Localize("tg.error.updatelang", nil)))
		return
	}

	this.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.clbk.lang.ru", nil)))
	// Remove loading animation
	this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, ""))
}

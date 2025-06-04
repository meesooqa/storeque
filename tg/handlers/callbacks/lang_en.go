package callbacks

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/tg-star-shop-bot/common/app"
	"github.com/meesooqa/tg-star-shop-bot/common/domain"
	"github.com/meesooqa/tg-star-shop-bot/common/lang"
	"github.com/meesooqa/tg-star-shop-bot/service/userservice"
)

type LangEnHandler struct {
	BaseHandler
	userService *userservice.Service
}

func NewLangEnHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *LangEnHandler {
	return &LangEnHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

func (this *LangEnHandler) GetData() string {
	return fmt.Sprintf("lang-%s", domain.UserSettingsLangValueEn)
}

func (this *LangEnHandler) Handle(ctx context.Context, loc lang.Localization, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID

	err := this.userService.SetChatLang(ctx, chatID, domain.UserSettingsLangValueEn)
	loc.SetLang(domain.UserSettingsLangValueEn)
	if err != nil {
		this.appDeps.Logger().Error("LangEnHandler", slog.Any("error", err))
		// Remove loading animation and show popup message
		this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, loc.Localize("tg.error.updatelang", nil)))
		return
	}

	this.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.clbk.lang.en", nil)))
	// Remove loading animation
	this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, ""))
}

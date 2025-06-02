package callbacks

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/common/domain"
	"tg-star-shop-bot-001/service/userservice"
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

func (this *LangEnHandler) Handle(ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID

	err := this.userService.SetChatLang(ctx, chatID, domain.UserSettingsLangValueEn)
	if err != nil {
		this.appDeps.Logger().Error("LangEnHandler", slog.Any("error", err))
		// Remove loading animation and show popup message
		this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, this.appDeps.Lang().Localize("tg.error.updatelang", nil)))
		return
	}

	this.bot.Send(tgbotapi.NewMessage(chatID, this.appDeps.Lang().Localize("tg.clbk.lang.en", nil)))
	// Remove loading animation
	this.bot.Send(tgbotapi.NewCallback(callbackQuery.ID, ""))
}

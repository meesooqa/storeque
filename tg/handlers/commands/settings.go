package commands

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/tg-star-shop-bot/common/app"
	"github.com/meesooqa/tg-star-shop-bot/common/lang"
	"github.com/meesooqa/tg-star-shop-bot/service/userservice"
)

type SettingsHandler struct {
	BaseHandler
	userService *userservice.Service
}

func NewSettingsHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *SettingsHandler {
	return &SettingsHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

func (this SettingsHandler) GetName() string {
	return "settings"
}

func (this SettingsHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", this.GetName()), nil)
}

func (this SettingsHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID

	userSettings, err := this.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		this.appDeps.Logger().Error("cmdSettings-GetUserSettings", slog.Any("error", err))
		this.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.error.getusersettings", nil)))
		return
	}

	text := fmt.Sprintf("*%s*:\n", loc.Localize("tg.cmd.settings.title", nil))
	text += fmt.Sprintf("• %s — *%s*\n", loc.Localize("tg.cmd.settings.lang", nil), userSettings.Lang)
	text += fmt.Sprintf("• %s — *%s*\n", loc.Localize("tg.cmd.settings.role", nil), userSettings.Role.Code)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ParseMode = "Markdown"
	this.bot.Send(msg)
}

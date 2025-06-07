package commands

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// SettingsHandler handles the /settings command in Telegram.
type SettingsHandler struct {
	BaseHandler
	userService userservice.UserService
}

// NewSettingsHandler creates a new instance of SettingsHandler.
func NewSettingsHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) *SettingsHandler {
	return &SettingsHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

// GetName returns the name of the command handler.
func (o *SettingsHandler) GetName() string {
	return "settings"
}

// GetDescription returns the localized description of the command handler
func (o *SettingsHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

// Handle processes the /settings command.
func (o *SettingsHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID

	userSettings, err := o.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		o.appDeps.Logger().Error("cmdSettings-GetUserSettings", slog.Any("error", err))
		o.bot.Send(tgbotapi.NewMessage(chatID, loc.Localize("tg.error.getusersettings", nil))) // nolint
		return
	}

	text := fmt.Sprintf("*%s*:\n", loc.Localize("tg.cmd.settings.title", nil))
	text += fmt.Sprintf("• %s — *%s*\n", loc.Localize("tg.cmd.settings.lang", nil), userSettings.Lang)
	text += fmt.Sprintf("• %s — *%s*\n", loc.Localize("tg.cmd.settings.role", nil), userSettings.Role.Code)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ParseMode = "Markdown"
	o.bot.Send(msg) // nolint
}

package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// StartHandler handles the /start command in Telegram
type StartHandler struct {
	BaseHandler
	userService userservice.UserService
}

// NewStartHandler creates a new instance of StartHandler
func NewStartHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) *StartHandler {
	return &StartHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

// GetName returns the name of the command handler
func (o *StartHandler) GetName() string {
	return "start"
}

// GetDescription returns the description of the command handler
func (o *StartHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

// Handle processes the /start command
func (o *StartHandler) Handle(_ context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	userName := inputMessage.From.FirstName
	if userName == "" {
		userName = inputMessage.From.LastName
	}
	if userName == "" {
		userName = "User" // TODO configurable default name
	}
	welcomeText := loc.Localize("tg.cmd.start.welcome", map[string]string{
		"UserName": userName,
	})
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, welcomeText)
	o.bot.Send(msg) // nolint

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇷🇺", "lang-ru"),
			tgbotapi.NewInlineKeyboardButtonData("🇺🇸", "lang-en"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇦", "lang-ru"),
			tgbotapi.NewInlineKeyboardButtonData("🇬🇧", "lang-en"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇧🇾", "lang-ru"),
			tgbotapi.NewInlineKeyboardButtonData("🇦🇺", "lang-en"),
		),
	)
	msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "Language:")
	msg.ReplyMarkup = keyboard
	o.bot.Send(msg) // nolint
}

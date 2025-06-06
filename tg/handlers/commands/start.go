package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

type StartHandler struct {
	BaseHandler
	userService userservice.UserService
}

func NewStartHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService userservice.UserService) *StartHandler {
	return &StartHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

func (o StartHandler) GetName() string {
	return "start"
}

func (o StartHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *StartHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	var userName string
	if inputMessage.From.FirstName != "" {
		userName = inputMessage.From.FirstName
	} else if inputMessage.From.LastName != "" {
		userName = inputMessage.From.LastName
	} else {
		userName = "User" // TODO configurable default name
	}
	userName = inputMessage.From.FirstName
	welcomeText := loc.Localize("tg.cmd.start.welcome", map[string]string{
		"UserName": userName,
	})
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, welcomeText)
	o.bot.Send(msg)

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
	o.bot.Send(msg)
}

package commands

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/common/domain"
	"tg-star-shop-bot-001/service/userservice"
)

type StartHandler struct {
	BaseHandler
	userService *userservice.Service
}

func NewStartHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *StartHandler {
	return &StartHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
		userService: userService,
	}
}

func (o *StartHandler) GetName() string {
	return "start"
}

func (o *StartHandler) GetDescription() string {
	return o.appDeps.Lang().Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *StartHandler) Handle(ctx context.Context, inputMessage *tgbotapi.Message) {
	// TODO langTag from user config
	//  "Welcome, {{.UserName}}!"
	//  "Use /help to see commands list."

	// TelegramID	2200662751
	// ChatID		5000386771
	user := &domain.User{
		ChatID: inputMessage.From.ID,
		// TelegramID: inputMessage.From.ID,
		Username:  inputMessage.From.UserName,
		FirstName: inputMessage.From.FirstName,
		LastName:  inputMessage.From.LastName,
	}

	// TODO err
	err := o.userService.Register(ctx, user)
	if err != nil {
		o.appDeps.Logger().Error("register error", slog.Any("error", err))
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, o.appDeps.Lang().Localize("tg.error.register", nil))
		o.bot.Send(msg)
		return
	}

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
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Language:")
	msg.ReplyMarkup = keyboard
	if _, err = o.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

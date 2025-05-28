package commands

import (
	"context"
	"fmt"
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

func NewStartHandler(appDeps *app.AppDeps, bot *tgbotapi.BotAPI, userService *userservice.Service) *StartHandler {
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
	return o.appDeps.Lang.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *StartHandler) Handle(inputMessage *tgbotapi.Message) {
	// TODO save user to DB
	// TODO langTag from user config
	//  "Welcome, {{.UserName}}!"
	//  "Use /help to see commands list."
	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	//o.bot.Send(msg)

	// TODO ctx := context.TODO()
	ctx := context.TODO()
	user := &domain.User{
		TelegramID: inputMessage.From.ID,
		Username:   inputMessage.From.UserName,
		FirstName:  inputMessage.From.FirstName,
		LastName:   inputMessage.From.LastName,
	}
	// TODO err
	err := o.userService.Register(ctx, user)
	if err != nil {
		o.appDeps.Logger.Error("register error", slog.Any("error", err))
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, o.appDeps.Lang.Localize("tg.error.register", nil))
		o.bot.Send(msg)
		return
	}
	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Добро пожаловать!")
	//o.bot.Send(msg)

	/*keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "btn1"),
			tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "btn2"),
		),
	)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите действие:")
	msg.ReplyMarkup = keyboard
	if _, err = o.bot.Send(msg); err != nil {
		log.Println(err)
	}*/
}

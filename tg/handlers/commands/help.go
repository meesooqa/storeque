package commands

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
)

type HelpHandler struct {
	BaseHandler
	availableCommands []CommandHandler
}

func NewHelpHandler(bot *tgbotapi.BotAPI, appDeps *app.AppDeps) *HelpHandler {
	return &HelpHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o *HelpHandler) GetName() string {
	return "help"
}

func (o *HelpHandler) GetDescription() string {
	return o.appDeps.Lang.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *HelpHandler) SetAvailableCommands(availableCommands []CommandHandler) {
	o.availableCommands = availableCommands
}

func (o *HelpHandler) Handle(inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	text := "*Справка по использованию бота*"
	text += "\n\n"
	text += "*Доступные команды:*"
	text += "\n\n"

	for _, cmd := range o.availableCommands {
		cmdHelpLine := fmt.Sprintf("• /%s — %s\n", cmd.GetName(), cmd.GetDescription())
		text += cmdHelpLine
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ParseMode = "Markdown"
	o.bot.Send(msg)
}

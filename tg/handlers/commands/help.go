package commands

import (
	"context"
	"fmt"
	"slices"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/common/lang"
)

type HelpHandler struct {
	BaseHandler
	allowedCommands []string
	commands        []CommandHandler
}

func NewHelpHandler(appDeps app.App, bot *tgbotapi.BotAPI) *HelpHandler {
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

func (o *HelpHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *HelpHandler) SetCommands(commands []CommandHandler) {
	o.commands = commands
}

func (o *HelpHandler) SetAllowedCommands(allowedCommands []string) {
	o.allowedCommands = allowedCommands
}

func (o *HelpHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	text := "*Справка по использованию бота*"
	text += "\n\n"
	text += "*Доступные команды:*"
	text += "\n\n"

	for _, cmd := range o.commands {
		if !slices.Contains(o.allowedCommands, cmd.GetName()) {
			continue
		}
		cmdHelpLine := fmt.Sprintf("• /%s — %s\n", cmd.GetName(), cmd.GetDescription(loc))
		text += cmdHelpLine
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ParseMode = "Markdown"
	o.bot.Send(msg)
}

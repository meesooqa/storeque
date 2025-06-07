package commands

import (
	"context"
	"fmt"
	"slices"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

// HelpHandler provides a handler for the /help command in a Telegram bot
type HelpHandler struct {
	BaseHandler
	allowedCommands []string
	commands        []CommandHandler
}

// NewHelpHandler creates a new instance of HelpHandler
func NewHelpHandler(appDeps app.App, bot *tgbotapi.BotAPI) *HelpHandler {
	return &HelpHandler{
		BaseHandler: BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

// GetName returns the name of the command handler
func (o *HelpHandler) GetName() string {
	return "help"
}

// GetDescription returns the description of the command handler
func (o *HelpHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

// SetCommands sets the list of command handlers that this HelpHandler will use
func (o *HelpHandler) SetCommands(commands []CommandHandler) {
	o.commands = commands
}

// SetAllowedCommands sets the list of allowed commands for this HelpHandler
func (o *HelpHandler) SetAllowedCommands(allowedCommands []string) {
	o.allowedCommands = allowedCommands
}

// Handle processes the /help command and sends a message with the bot's usage instructions
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
	o.bot.Send(msg) // nolint
}

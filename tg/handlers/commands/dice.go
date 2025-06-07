package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

// DiceHandler handles the /dice command in the Telegram bot
type DiceHandler struct {
	BaseHandler
}

// NewDiceHandler creates a new instance of DiceHandler with the provided app dependencies and bot API
func NewDiceHandler(appDeps app.App, bot *tgbotapi.BotAPI) *DiceHandler {
	return &DiceHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

// GetName returns the name of the command handler
func (o *DiceHandler) GetName() string {
	return "dice"
}

// GetDescription returns the localized description of the command handler
func (o *DiceHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

// Handle processes the /dice command, sending a dice emoji and returning the rolled value
func (o *DiceHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewDice(inputMessage.Chat.ID) // 1-6
	// msg := tgbotapi.NewDiceWithEmoji(inputMessage.Chat.ID, "🎰") // 1-64
	msg.ReplyParameters.MessageID = inputMessage.MessageID

	sentMsg, err := o.bot.Send(msg)
	if err != nil {
		m := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		o.bot.Send(m) // nolint
		return
	}
	m := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Value: %d", sentMsg.Dice.Value))
	m.ReplyParameters.MessageID = sentMsg.MessageID
	o.bot.Send(m) // nolint
	// fmt.Printf("Message sent successfully: %+v", sentMsg)
}

package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/tg-star-shop-bot/common/app"
	"github.com/meesooqa/tg-star-shop-bot/common/lang"
)

type DiceHandler struct {
	BaseHandler
}

func NewDiceHandler(appDeps app.App, bot *tgbotapi.BotAPI) *DiceHandler {
	return &DiceHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o DiceHandler) GetName() string {
	return "dice"
}

func (o DiceHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *DiceHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewDice(inputMessage.Chat.ID) // 1-6
	// msg := tgbotapi.NewDiceWithEmoji(inputMessage.Chat.ID, "🎰") // 1-64
	msg.ReplyParameters.MessageID = inputMessage.MessageID

	sentMsg, err := o.bot.Send(msg)
	if err != nil {
		m := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		o.bot.Send(m)
		return
	} else {
		m := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Value: %d", sentMsg.Dice.Value))
		m.ReplyParameters.MessageID = sentMsg.MessageID
		o.bot.Send(m)
	}
	// fmt.Printf("Message sent successfully: %+v", sentMsg)
}

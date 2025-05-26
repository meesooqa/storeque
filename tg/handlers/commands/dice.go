package commands

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
)

type DiceHandler struct {
	BaseHandler
}

func NewDiceHandler(bot *tgbotapi.BotAPI, appDeps *app.AppDeps) *DiceHandler {
	return &DiceHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o *DiceHandler) GetName() string {
	return "dice"
}

func (o *DiceHandler) GetDescription() string {
	return o.appDeps.Lang.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *DiceHandler) Handle(inputMessage *tgbotapi.Message) {
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

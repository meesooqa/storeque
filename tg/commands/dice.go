package commands

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

func (o *Commander) Dice(inputMessage *tgbotapi.Message) {
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

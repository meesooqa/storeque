package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (o *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	o.bot.Send(msg)
}

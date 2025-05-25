package commands

import tgbotapi "github.com/OvyFlash/telegram-bot-api"

func (o *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyParameters.MessageID = inputMessage.MessageID
	o.bot.Send(msg)
}

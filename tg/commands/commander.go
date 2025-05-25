package commands

import (
	"log"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (o *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "start":
		o.Start(update.Message)
	case "test":
		o.Tst(update.Message)
	default:
		o.Default(update.Message)
	}
}

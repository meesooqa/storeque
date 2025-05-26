package main

import (
	"log"
	"os"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"

	"tg-star-shop-bot-001/tg/commands"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	// bot, err := tgbotapi.NewBotAPI(token)
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "https://api.telegram.org/bot%s/test/%s")
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	commander := commands.NewCommander(bot)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			commander.HandleUpdate(update)
		}
	}
}

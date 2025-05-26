package main

import (
	"log"
	"os"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/tg/handlers"
)

func main() {
	appDeps := app.NewAppDeps()

	godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	// TODO bot, err := tgbotapi.NewBotAPI(token)
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "https://api.telegram.org/bot%s/test/%s")
	if err != nil {
		log.Fatal(err)
	}
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	handler := handlers.NewTelegramHandler(bot, appDeps)
	for update := range updates {
		handler.HandleUpdate(update)
	}
}

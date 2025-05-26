package main

import (
	"log"
	"log/slog"
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
	apiEndpoint := os.Getenv("TELEGRAM_API_ENDPOINT")

	var err error
	var bot *tgbotapi.BotAPI
	if apiEndpoint == "" {
		bot, err = tgbotapi.NewBotAPI(token)
	} else {
		bot, err = tgbotapi.NewBotAPIWithAPIEndpoint(token, apiEndpoint)
	}
	if err != nil {
		// appDeps.Logger.Error("NewBot", slog.Any("err", err))
		log.Fatal(err)
	}
	// bot.Debug = true
	appDeps.Logger.Info("Authorized", slog.String("Account", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	handler := handlers.NewTelegramHandler(bot, appDeps)
	for update := range updates {
		handler.HandleUpdate(update)
	}
}

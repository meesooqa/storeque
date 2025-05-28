package main

import (
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/db/repo"
	"tg-star-shop-bot-001/service/userservice"
	"tg-star-shop-bot-001/tg/handlers"
)

func main() {
	appDeps := app.NewAppDeps()

	db, err := appDeps.DBProvider.OpenDB()
	if err != nil {
		appDeps.Logger.Error("db opening error", slog.Any("error", err))
	}
	// TODO defer db.Close()
	userRepo := repo.NewUserRepo(db)
	userService := userservice.NewService(userRepo)

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	apiEndpoint := os.Getenv("TELEGRAM_API_ENDPOINT")

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

	bot.Debug = true

	// "from":{"id":5000386771,"is_bot":false,"first_name":"Stepan","last_name":"Test","language_code":"ru"}
	appDeps.Logger.Info("Authorized", slog.String("Account", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	handler := handlers.NewTelegramHandler(appDeps, bot, userService)
	for update := range updates {
		handler.HandleUpdate(update)
	}
}

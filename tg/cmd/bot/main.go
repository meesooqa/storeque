package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/db/repositories"
	"tg-star-shop-bot-001/service/userservice"
	"tg-star-shop-bot-001/tg/handlers"
)

func main() {
	appDeps := app.GetInstance()

	// TODO context: 10s crashes
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	ctx := context.Background()

	db, err := appDeps.DBProvider().Connect()
	if err != nil {
		appDeps.Logger().Error("db connection error", slog.Any("error", err))
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userSettingsRepo := repositories.NewUserSettingsRepository(db)
	userService := userservice.NewService(userRepo, userSettingsRepo)

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

	//bot.Debug = true

	appDeps.Logger().Info("Authorized", slog.String("Account", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	handler := handlers.NewTelegramHandler(appDeps, bot, userService)
	for update := range updates {
		handler.HandleUpdate(ctx, &update)
	}
}

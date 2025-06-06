package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/db/repositories"
	"github.com/meesooqa/storeque/service/locservice"
	"github.com/meesooqa/storeque/service/roleservice"
	"github.com/meesooqa/storeque/service/userservice"
	"github.com/meesooqa/storeque/tg/handlers"
	"github.com/meesooqa/storeque/tg/handlers/middleware"
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
	locService := locservice.NewService(appDeps, userService)
	roleService := roleservice.NewService()

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

	updatePreHandlers := []middleware.UpdatePreHandler{
		middleware.NewRegister(userService),
	}
	commandRouter := handlers.NewCommandRouter(appDeps, bot, locService, userService, roleService)
	callbackRouter := handlers.NewCallbackRouter(appDeps, bot, locService, userService)
	handler := handlers.NewTelegramHandler(appDeps, bot, updatePreHandlers, commandRouter, callbackRouter)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		handler.HandleUpdate(ctx, &update)
	}
}

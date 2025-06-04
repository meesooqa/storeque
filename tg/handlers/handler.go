package handlers

import (
	"context"
	"log/slog"
	"slices"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
	"github.com/meesooqa/storeque/tg/handlers/callbacks"
	"github.com/meesooqa/storeque/tg/handlers/commands"
)

type TelegramHandler struct {
	commands    map[string]commands.CommandHandler
	callbacks   map[string]callbacks.CallbackHandler
	userService *userservice.Service
	bot         *tgbotapi.BotAPI
	appDeps     app.App
}

func NewTelegramHandler(appDeps app.App, bot *tgbotapi.BotAPI, userService *userservice.Service) *TelegramHandler {
	return &TelegramHandler{
		commands:    commands.GetAll(appDeps, bot, userService),
		callbacks:   callbacks.GetAll(appDeps, bot, userService),
		userService: userService,
		bot:         bot,
		appDeps:     appDeps,
	}
}

func (this TelegramHandler) HandleUpdate(ctx context.Context, update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			this.appDeps.Logger().Info("recovered from panic", slog.Any("panicValue", panicValue))
		}
	}()

	loc := this.obtainLoc(ctx, update)
	allowedCommands, err := this.obtainAllowedCommands(ctx, update)
	if err != nil {
		this.appDeps.Logger().Error("HandleUpdate-allowedCommands", slog.Any("error", err))
		return
	}

	//if !CheckAuthorizedUser(update.Message.From.ID) {
	//	return
	//}

	// TODO SuccessfulPayment
	/*if update.Message.SuccessfulPayment != nil {
		handleSuccessfulPayment(update.Message)
		return
	}*/
	if update.CallbackQuery != nil {
		this.appDeps.Logger().Debug("clbk data", slog.String("data", update.CallbackQuery.Data))
		if callback, ok := this.callbacks[update.CallbackQuery.Data]; ok {
			callback.Handle(ctx, loc, update.CallbackQuery)
		} else {
			clbkHandler := callbacks.NewDefaultHandler(this.appDeps, this.bot)
			clbkHandler.Handle(ctx, loc, update.CallbackQuery)
		}
		return
	}

	if update.Message == nil {
		return
	}

	commandName := update.Message.Command()
	if command, ok := this.commands[commandName]; ok {
		if slices.Contains(allowedCommands, commandName) {
			if commandName == "help" {
				command.(*commands.HelpHandler).SetAllowedCommands(allowedCommands)
			}
			command.Handle(ctx, loc, update.Message)
		}
	} else {
		cmdHandler := commands.NewDefaultHandler(this.appDeps, this.bot)
		cmdHandler.Handle(ctx, loc, update.Message)
	}
}

func (this TelegramHandler) chatIdFromUpdate(update *tgbotapi.Update) int64 {
	var chatID int64 = 0
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	}
	return chatID
}

func (this TelegramHandler) obtainLoc(ctx context.Context, update *tgbotapi.Update) lang.Localization {
	chatID := this.chatIdFromUpdate(update)
	if chatID == 0 {
		return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), this.appDeps.Config().System.DefaultLangTag)
	}

	userSettings, err := this.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		this.appDeps.Logger().Error("TelegramHandler-GetUserSettings", slog.Any("error", err))
		return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), this.appDeps.Config().System.DefaultLangTag)
	}

	return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), userSettings.Lang)
}

func (this TelegramHandler) obtainAllowedCommands(ctx context.Context, update *tgbotapi.Update) ([]string, error) {
	chatID := this.chatIdFromUpdate(update)
	if chatID == 0 {
		return nil, nil
	}
	allowedCommands, err := this.userService.GetUserAllowedCommands(ctx, chatID)
	if err != nil {
		return nil, err
	}
	return allowedCommands, nil
}

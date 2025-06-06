package handlers

import (
	"context"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/service/locservice"
	"github.com/meesooqa/storeque/service/roleservice"
	"github.com/meesooqa/storeque/service/userservice"
	"github.com/meesooqa/storeque/tg/handlers/commands"
)

type CommandRouter struct {
	list           map[string]commands.CommandHandler
	defaultHandler commands.CommandHandler

	locService  locservice.LocService
	userService userservice.UserService
	roleService roleservice.RoleService
}

func NewCommandRouter(appDeps app.App, bot *tgbotapi.BotAPI, locService locservice.LocService, userService userservice.UserService, roleService roleservice.RoleService) *CommandRouter {
	return &CommandRouter{
		list:           commands.GetAll(appDeps, bot, userService),
		defaultHandler: commands.NewDefaultHandler(appDeps, bot),

		locService:  locService,
		userService: userService,
		roleService: roleService,
	}
}

func (this *CommandRouter) Route(ctx context.Context, update *tgbotapi.Update) error {
	if update.Message == nil || update.Message.Text == "" {
		return nil
	}

	chatID := this.chatIdFromUpdate(update)
	loc := this.locService.GetLoc(ctx, chatID)

	name := update.Message.Command()
	if name == "" {
		// not a command
		this.defaultHandler.Handle(ctx, loc, update.Message)
		return nil
	}

	handler, exists := this.list[name]
	if !exists {
		// command not found
		this.defaultHandler.Handle(ctx, loc, update.Message)
		return nil
	}

	userSettings, err := this.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		return err
	}
	allowedCommands := this.roleService.GetAllowedCommands(ctx, userSettings.Role.Code)

	if this.roleService.IsCommandAllowed(ctx, userSettings.Role.Code, name) {
		if name == "help" {
			handler.(*commands.HelpHandler).SetAllowedCommands(allowedCommands)
		}
		handler.Handle(ctx, loc, update.Message)
	} else {
		// TODO tg.error.commandnotallowed
	}
	return nil
}

func (this *CommandRouter) chatIdFromUpdate(update *tgbotapi.Update) int64 {
	var chatID int64 = 0
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	}
	return chatID
}

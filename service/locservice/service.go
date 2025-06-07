package locservice

import (
	"context"
	"log/slog"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

// Service implements the LocService interface
type Service struct {
	appDeps     app.App
	userService userservice.UserService
}

// NewService creates a new instance of Service with the provided dependencies
func NewService(appDeps app.App, userService userservice.UserService) *Service {
	return &Service{appDeps: appDeps, userService: userService}
}

// GetLoc retrieves the localization settings for a given chat ID
func (o *Service) GetLoc(ctx context.Context, chatID int64) lang.Localization {
	if chatID == 0 {
		return lang.NewUserLang(o.appDeps.Logger(), o.appDeps.LangBundle(), o.appDeps.Config().System.DefaultLangTag)
	}

	userSettings, err := o.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		o.appDeps.Logger().Error("TelegramHandler-GetUserSettings", slog.Any("error", err))
		return lang.NewUserLang(o.appDeps.Logger(), o.appDeps.LangBundle(), o.appDeps.Config().System.DefaultLangTag)
	}
	if userSettings == nil || userSettings.Lang == "" {
		return lang.NewUserLang(o.appDeps.Logger(), o.appDeps.LangBundle(), o.appDeps.Config().System.DefaultLangTag)
	}

	return lang.NewUserLang(o.appDeps.Logger(), o.appDeps.LangBundle(), userSettings.Lang)
}

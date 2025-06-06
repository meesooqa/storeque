package locservice

import (
	"context"
	"log/slog"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/service/userservice"
)

type Service struct {
	appDeps     app.App
	userService userservice.UserService
}

func NewService(appDeps app.App, userService userservice.UserService) *Service {
	return &Service{appDeps: appDeps, userService: userService}
}

func (this *Service) GetLoc(ctx context.Context, chatID int64) lang.Localization {
	if chatID == 0 {
		return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), this.appDeps.Config().System.DefaultLangTag)
	}

	userSettings, err := this.userService.GetUserSettings(ctx, chatID)
	if err != nil {
		this.appDeps.Logger().Error("TelegramHandler-GetUserSettings", slog.Any("error", err))
		return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), this.appDeps.Config().System.DefaultLangTag)
	}
	if userSettings == nil || userSettings.Lang == "" {
		return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), this.appDeps.Config().System.DefaultLangTag)
	}

	return lang.NewUserLang(this.appDeps.Logger(), this.appDeps.LangBundle(), userSettings.Lang)
}

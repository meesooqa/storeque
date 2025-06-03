package lang

import (
	"encoding/json"
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var moduleRegistrators []func(*i18n.Bundle)

func RegisterModuleTranslations(reg func(*i18n.Bundle)) {
	moduleRegistrators = append(moduleRegistrators, reg)
}

type UserLang struct {
	logger    *slog.Logger
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func NewBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, reg := range moduleRegistrators {
		reg(bundle)
	}
	return bundle
}

func NewUserLang(logger *slog.Logger, bundle *i18n.Bundle, userLangTag string) Localization {
	return &UserLang{
		logger:    logger,
		bundle:    bundle,
		localizer: i18n.NewLocalizer(bundle, userLangTag),
	}
}

func (this *UserLang) SetLang(userLangTag string) Localization {
	this.localizer = i18n.NewLocalizer(this.bundle, userLangTag)
	return this
}

func (this *UserLang) Localize(id string, tmplData map[string]string) string {
	text, err := this.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: tmplData,
	})
	if err != nil {
		this.logger.Error("i18n error", slog.Any("error", err))
		text = id
	}
	return text
}

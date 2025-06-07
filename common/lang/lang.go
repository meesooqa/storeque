package lang

import (
	"encoding/json"
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var moduleRegistrators []func(*i18n.Bundle)

// RegisterModuleTranslations registers a function that will be called to register translations for a specific module
func RegisterModuleTranslations(reg func(*i18n.Bundle)) {
	moduleRegistrators = append(moduleRegistrators, reg)
}

// UserLang implements the Localization interface for a specific user language
type UserLang struct {
	logger    *slog.Logger
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// NewBundle creates a new i18n.Bundle with the default language set to English.
func NewBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, reg := range moduleRegistrators {
		reg(bundle)
	}
	return bundle
}

// NewUserLang creates a new UserLang instance with the specified fields
func NewUserLang(logger *slog.Logger, bundle *i18n.Bundle, userLangTag string) Localization {
	return &UserLang{
		logger:    logger,
		bundle:    bundle,
		localizer: i18n.NewLocalizer(bundle, userLangTag),
	}
}

// SetLang sets the language for the UserLang instance and returns the updated Localization interface
func (o *UserLang) SetLang(userLangTag string) Localization {
	o.localizer = i18n.NewLocalizer(o.bundle, userLangTag)
	return o
}

// Localize localizes a message ID with optional template data for the UserLang instance
func (o *UserLang) Localize(id string, tmplData map[string]string) string {
	text, err := o.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: tmplData,
	})
	if err != nil {
		o.logger.Error("i18n error", slog.Any("error", err))
		text = id
	}
	return text
}

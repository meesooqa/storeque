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

type Localization interface {
	Localize(id string, tmplData map[string]string) string
}

type Lang struct {
	logger    *slog.Logger
	localizer *i18n.Localizer
}

func NewLang(logger *slog.Logger, langTag string) *Lang {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, reg := range moduleRegistrators {
		reg(bundle)
	}

	return &Lang{
		logger:    logger,
		localizer: i18n.NewLocalizer(bundle, langTag),
	}
}

func (o *Lang) Localize(id string, tmplData map[string]string) string {
	text, err := o.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: tmplData,
	})
	if err != nil {
		o.logger.Error("i18n error", slog.Any("err", err))
		text = id
	}
	return text
}

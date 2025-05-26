package lang

import (
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localization interface {
	Localize(id string, tmplData map[string]string) string
}

type Lang struct {
	localizer *i18n.Localizer
}

func NewLang(langTag string) *Lang {
	return &Lang{
		localizer: i18n.NewLocalizer(initI18n(), langTag),
	}
}

func (o *Lang) Localize(id string, tmplData map[string]string) string {
	text, err := o.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: tmplData,
	})
	if err != nil {
		log.Println("i18n error:", err)
		text = id
	}
	return text
}

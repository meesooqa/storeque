package tg

import (
	"embed"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

//go:embed locales/*.json
var localeFS embed.FS

// LoadLocales loads translation files from the embedded filesystem into the provided i18n bundle
func LoadLocales(bundle *i18n.Bundle) error {
	files, err := localeFS.ReadDir("locales")
	if err != nil {
		return err
	}
	for _, f := range files {
		_, err = bundle.LoadMessageFileFS(localeFS, fmt.Sprintf("locales/%s", f.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

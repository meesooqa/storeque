package tg

import (
	"embed"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/meesooqa/tg-star-shop-bot/common/lang"
)

//go:embed locales/*.json
var localeFS embed.FS

func init() {
	lang.RegisterModuleTranslations(func(bundle *i18n.Bundle) {
		files, _ := localeFS.ReadDir("locales")
		for _, f := range files {
			_, _ = bundle.LoadMessageFileFS(localeFS, fmt.Sprintf("locales/%s", f.Name()))
		}
	})
}

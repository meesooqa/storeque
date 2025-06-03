package lang

type Localization interface {
	SetLang(userLangTag string) Localization
	Localize(id string, tmplData map[string]string) string
}

package lang

// Localization interface defines methods for localization
type Localization interface {
	SetLang(userLangTag string) Localization
	Localize(id string, tmplData map[string]string) string
}

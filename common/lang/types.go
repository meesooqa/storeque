package lang

type Localization interface {
	Localize(id string, tmplData map[string]string) string
}

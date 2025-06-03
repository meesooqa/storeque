package lang

import (
	"bytes"
	"io"
	"log/slog"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func TestRegisterModuleTranslations_CallsRegistrators(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	called := false
	RegisterModuleTranslations(func(b *i18n.Bundle) {
		called = true
	})

	// Trigger registration
	NewBundle()

	assert.True(t, called)
}

func TestNewBundle_RegistersJSONUnmarshal(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	bundle := NewBundle()

	msgJSON := `[{
		"test": "Test Message",
		"sub": {
			"key1": "value1",
			"sub1": {
				"key2": "value2"
			}
		}
	}]`

	_, err := bundle.ParseMessageFileBytes([]byte(msgJSON), "test.en.json")
	require.NoError(t, err)

	localizer := i18n.NewLocalizer(bundle, language.English.String())

	res, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: "test"})
	require.NoError(t, err)
	assert.Equal(t, "Test Message", res)

	res, err = localizer.Localize(&i18n.LocalizeConfig{MessageID: "sub.key1"})
	require.NoError(t, err)
	assert.Equal(t, "value1", res)

	res, err = localizer.Localize(&i18n.LocalizeConfig{MessageID: "sub.sub1.key2"})
	require.NoError(t, err)
	assert.Equal(t, "value2", res)
}

func TestNewUserLang_InitializesWithProvidedLang(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	RegisterModuleTranslations(func(b *i18n.Bundle) {
		b.AddMessages(language.English, &i18n.Message{ID: "hello", Other: "Hello"})
		b.AddMessages(language.Russian, &i18n.Message{ID: "hello", Other: "Привет"})
	})

	bundle := NewBundle()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	userLang := NewUserLang(logger, bundle, "ru").(Localization)

	assert.Equal(t, "Привет", userLang.Localize("hello", nil))
}

func TestSetLang_ChangesLanguage(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	RegisterModuleTranslations(func(b *i18n.Bundle) {
		b.AddMessages(language.English, &i18n.Message{ID: "greeting", Other: "Hello"})
		b.AddMessages(language.Russian, &i18n.Message{ID: "greeting", Other: "Привет"})
	})

	bundle := NewBundle()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	userLang := NewUserLang(logger, bundle, "en").(Localization)

	assert.Equal(t, "Hello", userLang.Localize("greeting", nil))
	userLang.SetLang("ru")
	assert.Equal(t, "Привет", userLang.Localize("greeting", nil))
}

func TestLocalize_ReturnsIDOnError(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	RegisterModuleTranslations(func(b *i18n.Bundle) {
		// No messages added
	})

	bundle := NewBundle()
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	userLang := NewUserLang(logger, bundle, "en").(Localization)

	id := "unknown"
	result := userLang.Localize(id, nil)
	assert.Equal(t, id, result)
	assert.Contains(t, buf.String(), "i18n error")
	assert.Contains(t, buf.String(), id)
}

func TestLocalize_UsesTemplateData(t *testing.T) {
	oldRegs := moduleRegistrators
	defer func() { moduleRegistrators = oldRegs }()
	moduleRegistrators = nil

	RegisterModuleTranslations(func(b *i18n.Bundle) {
		b.AddMessages(language.English, &i18n.Message{
			ID:    "welcome",
			Other: "Hello {{.name}}",
		})
	})

	bundle := NewBundle()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	userLang := NewUserLang(logger, bundle, "en").(Localization)

	tmplData := map[string]string{"name": "Alice"}
	res := userLang.Localize("welcome", tmplData)
	assert.Equal(t, "Hello Alice", res)
}

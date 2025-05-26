package lang

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func initI18n() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	err := loadMessageFiles(bundle, "locales")
	if err != nil {
		log.Fatalf("i18n load error: %v", err)
	}
	return bundle
}

// loadAllMessageFiles goes through all *.json and *.yaml in the specified folder
func loadMessageFiles(bundle *i18n.Bundle, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext == ".json" || ext == ".yaml" || ext == ".yml" {
			if _, err := bundle.LoadMessageFile(path); err != nil {
				return err
			}
			// log.Printf("Loaded locale file: %s", path)
		}
		return nil
	})
}

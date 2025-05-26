package lang

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localeFS embed.FS

func initI18n() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	err := loadMessageFiles(bundle, localeFS)
	if err != nil {
		log.Fatalf("i18n load error: %v", err)
	}
	return bundle
}

// loadAllMessageFiles goes through all *.json in the specified folder
func loadMessageFiles(bundle *i18n.Bundle, fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".json" {
			file, err := fsys.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				return err
			}
			if _, err := bundle.ParseMessageFileBytes(data, path); err != nil {
				return err
			}
			log.Printf("Loaded locale file: %s", path)
		}
		return nil
	})
}

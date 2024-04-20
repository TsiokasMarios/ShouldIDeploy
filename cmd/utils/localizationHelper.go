package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
)

var bundle *i18n.Bundle

func LoadBundles() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load English translations
	enBytes, err := os.ReadFile("../../translations/en.json")
	if err != nil {
		log.Fatalf("Error loading English translations: %v", err)
	}
	_, err = bundle.ParseMessageFileBytes(enBytes, "../../translations/en.json")
	if err != nil {
		return
	}

	// Load Greek translations
	elBytes, err := os.ReadFile("../../translations/el.json")
	if err != nil {
		log.Fatalf("Error loading Greek translations: %v", err)
	}
	_, err = bundle.ParseMessageFileBytes(elBytes, "../../translations/el.json")

	if err != nil {
		return
	}

}

func Localize(toLocalize string, language string) string {
	println(language)
	localizer := i18n.NewLocalizer(bundle, language)
	println(bundle.LanguageTags())
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: toLocalize})
}

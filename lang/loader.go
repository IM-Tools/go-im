// Package lang is a part of the package carbon for i18n
package lang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// langDir language directory
	langDir = "./lang/"

	// invalidLocaleError returns an invalid locale error.
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale %q, please see the directory %q for all valid locales", locale, langDir)
	}

	// invalidJsonFileError returns an invalid json file error.
	invalidJsonFileError = func(file string) error {
		return fmt.Errorf("invalid json file %q, please make sure the json file is valid", langDir+file)
	}
)

// LoadLocale loads a given locale file.
func LoadLocale(locale string) (map[string]string, error) {
	resources := make(map[string]string)
	fileName := langDir + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return resources, invalidLocaleError(locale)
	}
	if err := json.Unmarshal(bytes, &resources); err != nil {
		return resources, invalidJsonFileError(fileName)
	}
	return resources, nil
}

package store

import (
	"fmt"
	"strings"
	logger "merchants_of_galaxy/logger"
)
var translations = make(map[string]string)


// GetTranslation returns a Translation from alien to roman numerals
func GetTranslation(input string) (string, error) {
	tokens := strings.Split(input, " ")
	output := ""

	for _, token := range tokens {
		translation, ok := translations[token]
		if ok {
			output = fmt.Sprintf("%s%s", output, translation)
		} else {
			logger.Error.Printf("Cannot translate %s to a roman numeral", token)
			return "", fmt.Errorf("Cannot translate %s to a roman numeral", token)
		}
	}

	return output, nil
}

// SaveTranslation saves alien to roman numeral translations
func SaveTranslation(alien string, roman string) {
	_, present := translations[alien]

	if !present {
		translations[alien] = roman
		logger.Trace.Printf("Saved translation %s = %s", alien, roman)
	} else {
		logger.Error.Printf("Translation for %s already exists", alien)
	}
}
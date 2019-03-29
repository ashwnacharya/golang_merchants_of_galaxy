package query

import (
	"regexp"
	"log"
	"fmt"
	"merchants_of_galaxy/commands"
	"merchants_of_galaxy/rns"
	"merchants_of_galaxy/store"
	logger "merchants_of_galaxy/logger"
)

type translator struct {

}

const translatorRegex = "how much is (?P<alien_amount>[a-zA-Z ]+) \\?"

func init() {
	var translatorCommand translator
	commands.Register("translator", translatorRegex, translatorCommand)
}

func (t translator) Execute(input string) (string) {
	var reg = regexp.MustCompile(translatorRegex)

	match := reg.FindStringSubmatch(input)

	result := make(map[string]string)
    for i, name := range reg.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
	}

	alienAmount := result["alien_amount"]

	romanAmount, translationErr := store.GetTranslation(alienAmount)
	if translationErr != nil {
		logger.Error.Printf("Cannot convert from alien numeral %s to valid roman numeral", alienAmount)
		return ""
	}

	amount, err := rns.ConvertToInteger(romanAmount)

	if err != nil {
		log.Fatal("Cannot convert from alien numeral to valid integer")
	}

	return fmt.Sprintf("%s is %d", alienAmount, amount)

}
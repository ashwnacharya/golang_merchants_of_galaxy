package query

import (
	"regexp"
	"fmt"
	"merchants_of_galaxy/commands"
	"merchants_of_galaxy/store"
	"merchants_of_galaxy/rns"
	logger "merchants_of_galaxy/logger"
)

type cost struct {

}

const costRegex = "how many Credits is (?P<alien_quantity>[a-zA-Z ]+) (?P<item>[a-zA-Z]+) \\?"

func init() {
	var costCommand cost
	commands.Register("cost", costRegex, costCommand)
}

func (c cost) Execute(input string) (string) {
	var reg = regexp.MustCompile(costRegex)

	match := reg.FindStringSubmatch(input)

	result := make(map[string]string)
    for i, name := range reg.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
	}

	alienQuantity := result["alien_quantity"]
	item := result["item"]

	romanQuantity, translationErr := store.GetTranslation(alienQuantity)
	if translationErr != nil {
		logger.Error.Printf("Cannot convert from alien numeral %s to valid roman numeral", alienQuantity)
	}

	quantity, err := rns.ConvertToInteger(romanQuantity)

	if err != nil {
		logger.Error.Println("Cannot convert from alien numeral to valid integer")
	}

	cost:= quantity * store.Prices[item]

	return fmt.Sprintf("%s %s is %d credits", alienQuantity, item, cost)
}
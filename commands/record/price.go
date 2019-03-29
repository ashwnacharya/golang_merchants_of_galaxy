package query

import (
	"fmt"
	"regexp"
	"log"
	"strconv"
	"merchants_of_galaxy/commands"
	"merchants_of_galaxy/store"
	"merchants_of_galaxy/rns"
	logger "merchants_of_galaxy/logger"
)

type price struct {

}

const priceRegex = "(?P<quantity_in_alien>[a-zA-Z ]+) (?P<item>[a-zA-Z]+) is (?P<amount>[0-9]+) Credits"

func init() {
	var priceCommand price
	commands.Register("price", priceRegex, priceCommand)
}

func (p price) Execute(input string) (string) {
	var reg = regexp.MustCompile(priceRegex)

	match := reg.FindStringSubmatch(input)

	result := make(map[string]string)
    for i, name := range reg.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
	}

	alienQuantity := result["quantity_in_alien"]
	romanQuantity,translationErr := store.GetTranslation(alienQuantity)

	if translationErr != nil {
		logger.Error.Printf("Cannot convert from alien numeral %s to valid roman numeral", alienQuantity)
		return ""
	}

	quantity, err := rns.ConvertToInteger(romanQuantity)

	if err != nil {
		log.Fatal("Cannot convert from alien numeral to valid integer")
	}

	item := result["item"]
	amount, _ := strconv.Atoi(result["amount"])

	store.Prices[item] = int(amount/quantity)

	fmt.Println(store.Prices)
	return ""
}
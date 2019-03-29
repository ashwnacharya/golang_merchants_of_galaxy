package query

import (
	"regexp"
	"merchants_of_galaxy/commands"
	"merchants_of_galaxy/store"
	logger "merchants_of_galaxy/logger"
)

type decimal struct {

}

const decimalRegex = "(?P<alien>[a-zA-Z ]+) is (?P<roman>[IVXLCDM])"

func init() {
	var decimalCommand decimal
	commands.Register("decimal", decimalRegex, decimalCommand)
}

func (d decimal) Execute(input string) (string) {
	var reg = regexp.MustCompile(decimalRegex)

	match := reg.FindStringSubmatch(input)

	result := make(map[string]string)
    for i, name := range reg.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
	}

	store.SaveTranslation(result["alien"], result["roman"])
	logger.Info.Printf("Saved translation %s = %s", result["alien"], result["roman"])
	return ""
}
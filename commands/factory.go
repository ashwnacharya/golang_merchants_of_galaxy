package commands

import (
	"regexp"
	logger "merchants_of_galaxy/logger"
)

// Command interface defines the methods that every command should implement
type Command interface {
	Execute(input string) (string)
}

var commandRegexes = make(map[string]string)
var commands = make(map[string]Command)

// Register method registers various commands that are available, to the command map.
func Register(commandType string, regex string, command Command) {
	logger.Trace.Printf("Registering command type %s", commandType)
	commandRegexes[commandType] = regex
	commands[commandType] = command
}


// Execute function detects the command type and executes it
func Execute(input string) (string) {
	for commandType, commandRegex := range commandRegexes {
		match, _ := regexp.MatchString(commandRegex, input)
		if match == true {
			logger.Trace.Printf("Command matched with command type %s", commandType)
			output:= commands[commandType].Execute(input)
			logger.Trace.Printf("Received output %s", output)
			return output
		}
	}

	logger.Trace.Printf("Command did not match with any command type")
	return "I have no idea what you are talking about"
}
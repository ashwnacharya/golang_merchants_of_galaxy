package main

import (
	"fmt"
	"os"
	"bufio"
	"merchants_of_galaxy/commands"
	_ "merchants_of_galaxy/commands/query"
	_ "merchants_of_galaxy/commands/record"
	logger "merchants_of_galaxy/logger"
)


func main() {

	logger.Trace.Println("Starting the program")

	for true {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		logger.Trace.Printf("Received input %s\n", input)

		if input == "exit\n" {
			logger.Trace.Println("Exiting the program")
			return
		}
		output := commands.Execute(input)
		fmt.Println(output)
	}
}
package commands

import (
	"fmt"
	"os"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func()
}

func getDefaultCommands() map[string]CLICommand {
	commands := map[string]CLICommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
	}

	return commands
}

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := getDefaultCommands()

	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
}

package commands

import (
	"fmt"
	"github.com/Kenedy228/pokedex/internal/requests"
	"os"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func() error
}

const mapURL = "https://pokeapi.co/api/v2/location/"

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
		"map": {
			Name:        "map",
			Description: "Displays location-areas",
			Callback:    commandMap,
		},
	}

	return commands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := getDefaultCommands()

	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}

	return nil
}

func commandMap() error {
	mapper, err := requests.NewMapper(mapURL)

	if err != nil {
		return err
	}

	locations, err := mapper.Handle()

	if err != nil {
		return err
	}

	for _, v := range locations {
		for _, a := range v.Areas {
			fmt.Printf("%s\n", a.Name)
		}
	}

	return nil
}

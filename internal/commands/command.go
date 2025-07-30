package commands

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Kenedy228/pokedex/internal/requests"
)

func GetDefaultCommands() map[string]Command {
	commands := map[string]Command{
		"exit":    NewExitCommand(),
		"help":    NewHelpCommand(),
		"map":     NewMapCommand(),
		"explore": NewExploreCommand(),
		"catch":   NewCatchCommand(),
	}

	return commands
}

func (c CLICommand) Name() string {
	return c.name
}

func (c CLICommand) Description() string {
	return c.description
}

func NewExitCommand() Command {
	exitCommand := ExitCommand{
		CLICommand: CLICommand{
			name:        "exit",
			description: "Exit the Pokedex",
		},
	}

	return exitCommand
}

func (cmd ExitCommand) Callback(m *requests.Mapper, args []string) error {
	err := cmd.ValidateArgs(args)

	if err != nil {
		return err
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (cmd ExitCommand) ValidateArgs(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("Too many arguments provided for the %s command", cmd.Name())
	}

	return nil
}

func NewHelpCommand() Command {
	helpCommand := HelpCommand{
		CLICommand: CLICommand{
			name:        "help",
			description: "Displays a help message",
		},
	}

	return helpCommand
}

func (cmd HelpCommand) Callback(m *requests.Mapper, args []string) error {
	err := cmd.ValidateArgs(args)

	if err != nil {
		return err
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	commands := GetDefaultCommands()

	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.Name(), value.Description())
	}

	return nil
}

func (cmd HelpCommand) ValidateArgs(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("Too many arguments provided for the %s command", cmd.Name())
	}

	return nil
}

func NewMapCommand() Command {
	mapCommand := MapCommand{
		CLICommand: CLICommand{
			name:        "map",
			description: "Displays location-areas",
		},
	}

	return mapCommand
}

func (cmd MapCommand) Callback(m *requests.Mapper, args []string) error {
	err := cmd.ValidateArgs(args)

	if err != nil {
		return err
	}

	areas, err := m.GetAreas()

	if err != nil {
		return err
	}

	output := ""

	for key, _ := range areas {
		output += key + " "
	}

	fmt.Println(output)

	return nil
}

func (cmd MapCommand) ValidateArgs(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("Too many arguments provided for the %s command", cmd.Name())
	}

	return nil
}

func NewExploreCommand() Command {
	exploreCommand := ExploreCommand{
		CLICommand: CLICommand{
			name:        "explore",
			description: "finds pokemons in area",
		},
	}

	return exploreCommand
}

func (cmd ExploreCommand) Callback(m *requests.Mapper, args []string) error {
	err := cmd.ValidateArgs(args)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[1])
	pokemons, err := m.GetPocemonsByArea(args[1])

	if err != nil {
		return err
	}

	if len(pokemons) == 0 {
		return fmt.Errorf("Not found pokemons in provided area")
	}

	fmt.Println("Found Pokemon:")

	for _, p := range pokemons {
		fmt.Printf("- %s\n", p.Name)
	}

	return nil
}

func (cmd ExploreCommand) ValidateArgs(args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("Too many arguments provided for the %s command", cmd.Name())
	}

	return nil
}

func NewCatchCommand() Command {
	catchCommand := CatchCommand{
		CLICommand: CLICommand{
			name:        "catch",
			description: "game catch pokemon",
		},
	}

	return catchCommand
}

func (cmd CatchCommand) Callback(m *requests.Mapper, args []string) error {
	stats, err := m.FindPokemonExperience(args[1])

	if err != nil {
		return err
	}

	randNum := rand.Float32()
	floatStats := float32(stats.Experience)

	fmt.Printf("Throwing a Pokeball at %s...\n", args[1])

	if randNum > floatStats {
		fmt.Printf("%s escaped!", args[1])
		return nil
	}

	fmt.Printf("%s was caught!\n", args[1])

	return nil
}

func (cmd CatchCommand) ValidateArgs(args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("Too many arguments provided for the %s command", cmd.Name())
	}

	return nil
}

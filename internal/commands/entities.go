package commands

import "github.com/Kenedy228/pokedex/internal/requests"

type Command interface {
	Name() string
	Description() string
	Callback(*requests.Mapper, []string) error
	ValidateArgs([]string) error
}

type CLICommand struct {
	name        string
	description string
}

type ExitCommand struct {
	CLICommand
}

type HelpCommand struct {
	CLICommand
}

type MapCommand struct {
	CLICommand
}

type ExploreCommand struct {
	CLICommand
}

type CatchCommand struct {
	CLICommand
}

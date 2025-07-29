package commands

import (
	"fmt"
	"github.com/Kenedy228/pokedex/common"
)

type CommandHandler struct {
	commands map[string]CLICommand
}

func NewHandler() CommandHandler {
	commands := getDefaultCommands()

	handler := CommandHandler{commands: commands}

	return handler
}

func (h CommandHandler) HandleCommand(userInput string) {
	cleaned := common.CleanInput(userInput)

	if len(cleaned) == 0 {
		fmt.Println("Provide command first")
		return
	}

	var err error

	switch cleaned[0] {
	case "exit":
		err = h.commands["exit"].Callback()
	case "help":
		err = h.commands["help"].Callback()
	case "map":
		err = h.commands["map"].Callback()
	default:
		fmt.Println("Unknown command")
	}

	if err != nil {
		fmt.Printf("%v", err)
	}
}

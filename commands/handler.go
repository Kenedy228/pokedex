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

	switch cleaned[0] {
	case "exit":
		h.commands["exit"].Callback()
	case "help":
		h.commands["help"].Callback()
	default:
		fmt.Println("Unknown command")
	}
}

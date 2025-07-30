package commands

import (
	"fmt"

	"github.com/Kenedy228/pokedex/internal/common"
	"github.com/Kenedy228/pokedex/internal/requests"
)

type CommandHandler struct {
	commands map[string]Command
	mapper   *requests.Mapper
}

func NewHandler() CommandHandler {
	commands := GetDefaultCommands()

	handler := CommandHandler{commands: commands}
	mapper, _ := requests.NewMapper()
	handler.mapper = mapper

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
		err = h.commands["exit"].Callback(h.mapper, cleaned)
	case "help":
		err = h.commands["help"].Callback(h.mapper, cleaned)
	case "map":
		err = h.commands["map"].Callback(h.mapper, cleaned)
	case "explore":
		err = h.commands["explore"].Callback(h.mapper, cleaned)
	case "catch":
		err = h.commands["catch"].Callback(h.mapper, cleaned)
	default:
		fmt.Println("Unknown command")
	}

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

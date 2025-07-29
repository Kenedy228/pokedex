package main

import (
	"bufio"
	"fmt"
	"github.com/Kenedy228/pokedex/internal/commands"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	handler := commands.NewHandler()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		handler.HandleCommand(scanner.Text())
	}
}

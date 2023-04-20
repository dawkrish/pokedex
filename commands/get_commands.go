package commands

import(
	"github.com/krishnanshagarwal112/pokedex/models"
)

func GetCommands() map[string]models.CliCommand {
	return map[string]models.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapB,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
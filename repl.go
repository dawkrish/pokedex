package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/krishnanshagarwal112/pokedex/commands"
	"github.com/krishnanshagarwal112/pokedex/models"
)


func startRepl(cfg *models.Config) {
	commands := commands.GetCommands()

	for {
		fmt.Printf("\nPokedex > ")
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		name := reader.Text()
		fmt.Printf("\n")
		if len(name) == 0 {
			fmt.Println("Enter a command !")
			continue
		}

		val, ok := commands[name]
		if ok {
			commands[val.Name].Callback(cfg)
		} else {
			fmt.Println("Unknown command !")
		}
	}
}


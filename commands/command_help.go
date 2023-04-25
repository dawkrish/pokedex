package commands

import (
	"fmt"
	"github.com/krishnanshagarwal112/pokedex/models"
	"errors"
)

func commandHelp(cfg *models.Config,s string) error {
	if len(s) > 0{
		return errors.New("just type `help`")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf(("Usage:\n\n"))

	for _, command := range GetCommands(){
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	
	return nil
}
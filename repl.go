package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/krishnanshagarwal112/pokedex/commands"
	"github.com/krishnanshagarwal112/pokedex/models"
)

func StartRepl(cfg *models.Config) {
	commands := commands.GetCommands()

	for {
		fmt.Printf("\nPokedex > ")
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		name := reader.Text()
		lower := strings.ToLower(name)
		arguments := strings.Fields(lower)
		fmt.Println(arguments)

		fmt.Printf("\n")
		if len(name) == 0 {
			fmt.Println("Enter a command !")
			continue
		}
		if len(arguments) == 1{
			val, ok := commands[name]
			if ok {
				err := commands[val.Name].Callback(cfg,"")
				if err != nil{
					fmt.Printf("%v\n",err)
					continue
				}
			} else {
				fmt.Println("Unknown command !")
			}
		}	
		if len(arguments) == 2{
			if arguments[0] == "explore"{
				err:=commands[arguments[0]].Callback(cfg,arguments[1])
				if err != nil{
					fmt.Printf("%v\n",err)
					continue
			}
		}
	}
}
}
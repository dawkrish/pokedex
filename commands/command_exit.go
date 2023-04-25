package commands

import (
	"os"
	"github.com/krishnanshagarwal112/pokedex/models"
	"errors"

)



func commandExit(cfg *models.Config, s string) error {
	if len(s) > 0{
		return errors.New("just type `exit`")
	}
	os.Exit(0)
	return nil
}
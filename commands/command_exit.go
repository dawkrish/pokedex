package commands

import (
	"os"
	"github.com/krishnanshagarwal112/pokedex/models"

)



func commandExit(cfg *models.Config, s string) error {
	os.Exit(0)
	return nil
}
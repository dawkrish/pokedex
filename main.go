package main

import (
	"github.com/krishnanshagarwal112/pokedex/models"
)


func main() {
	cfg := &models.Config{Prev: "",Next: "https://pokeapi.co/api/v2/location-area/"}
	startRepl(cfg)
}



package main

import (
	"time"

	"github.com/krishnanshagarwal112/pokedex/cache"
	"github.com/krishnanshagarwal112/pokedex/models"
)



func main() {
	cfg := &models.Config{
		Prev: "",
		Next: "https://pokeapi.co/api/v2/location-area/?",
		Caching: cache.NewCache(5*time.Second),}
	StartRepl(cfg)
}



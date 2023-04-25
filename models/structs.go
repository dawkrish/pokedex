package models

import (
	"sync"
	"time"
)

type CliCommand struct {
	Name       string
	Description string
	Callback    func(*Config, string) error
}

type Config struct{
	Prev string
	Next string
	Caching Cache
	Pokemons []string
}


type LocationResult struct{
	Prev string `json:"previous"`
	Next string	`json:"next"`
	Results []struct{
		Name string `json:"name"`
		Url string 	`json:"url"`
	} `json:"results"`
}

type ExploreResult struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Cache struct{
	Caches map[string]CacheEntry
	Mux *sync.Mutex
}

type CacheEntry struct{
	CreatedAt time.Time
	Val []byte
}

type CatchResult struct {
	BaseExperience int `json:"base_experience"`
}


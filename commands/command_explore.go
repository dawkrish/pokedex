package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/krishnanshagarwal112/pokedex/cache"
	"github.com/krishnanshagarwal112/pokedex/models"
)

func commandExplore(cfg *models.Config, argument string) error{
	if(argument == ""){
		return errors.New("enter a location too")
	}
	req_url := "https://pokeapi.co/api/v2/location-area/" +argument
	val,ok := cache.Get(&cfg.Caching,req_url)
	result := models.ExploreResult{}
	var body []byte

	fmt.Println("Found in Cache : ",ok)

	if ok{
		body = val
	}else{
		res, err := http.Get(req_url)
		fmt.Println("ERROR : ",err)
		if err != nil{
			return err
	}

		respBody, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return err
		}
		body = respBody
		cache.Add(&cfg.Caching,req_url,body)
	}

	err:= json.Unmarshal(body,&result)
	if err != nil{
		fmt.Println("WRONG JSON !")
		return err
	}
	pokemons := result.PokemonEncounters
	fmt.Printf("Exploring %v...\n",argument)
	fmt.Printf("Found Pokemon:\n")
	for _,pokemon := range pokemons{
		fmt.Println(" - ",pokemon.Pokemon.Name)
	}
	
	return nil
}
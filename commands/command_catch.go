package commands

import ("github.com/krishnanshagarwal112/pokedex/models"
		"github.com/krishnanshagarwal112/pokedex/cache"
		"errors"
		"net/http"
		"encoding/json"
		"fmt"
		"io"
		"math/rand")

func commandCatch(cfg *models.Config, argument string) error{
	if(argument == ""){
		return errors.New("enter a pokemon too")
	}
	req_url := "https://pokeapi.co/api/v2/pokemon/" +argument
	val,ok := cache.Get(&cfg.Caching,req_url)
	result := models.CatchResult{}
	var body []byte


	if ok{
		body = val
	}else{
		res, err := http.Get(req_url)
		if err != nil{
			fmt.Print("http return")
			return err
	}

		respBody, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			fmt.Print("http converting resp to byte error")
			return err
		}
		body = respBody
		cache.Add(&cfg.Caching,req_url,body)
	}

	err:= json.Unmarshal(body,&result)
	if err != nil{
		fmt.Print("json to struct error")
		return err
	}
	
	fmt.Println("Throwing a Pokeball at "+argument+"...")

	base_expeirence := result.BaseExperience
	random_exp := rand.Intn(base_expeirence)

	if random_exp > 40 {
		fmt.Printf("%s escaped!\n",argument)
		return nil
	}
	fmt.Printf("%s was caught!\n",argument)

	cfg.Pokemons = append(cfg.Pokemons, argument)
	return nil
}
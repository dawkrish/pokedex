package commands

import(
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"

	"github.com/krishnanshagarwal112/pokedex/models"
	"github.com/krishnanshagarwal112/pokedex/cache"
)

func commandMap(cfg *models.Config, s string) error {
	val,ok := cache.Get(&cfg.Caching,cfg.Next)
	result := models.LocationResult{}
	var body []byte
	
	if ok {
		body = val

	}else{
		res, err := http.Get(cfg.Next)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return err
		}
		body = respBody
		// add it to cache !
		cache.Add(&cfg.Caching,cfg.Next,body)
	}

	err := json.Unmarshal(body, &result)
		if err != nil {
		return err
	}
	cfg.Prev = result.Prev
	cfg.Next = result.Next

	areas := result.Results
	for _,area := range areas{
		fmt.Println(area.Name)
	}
	return err
}


func commandMapB(cfg *models.Config, s string) error {
	if(cfg.Prev == ""){
		err := errors.New("cannot go previous now ")
		return err
	}

	val,ok := cache.Get(&cfg.Caching,cfg.Prev)
	result := models.LocationResult{}
	var body []byte

	if ok {
		body = val

	}else{
		res, err := http.Get(cfg.Prev)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return err
		}
		body = respBody
		// add it to cache !
		cache.Add(&cfg.Caching,cfg.Prev,body)
	}

	err := json.Unmarshal(body, &result)
		if err != nil {
		return err
	}
	cfg.Prev = result.Prev
	cfg.Next = result.Next

	areas := result.Results
	for _,area := range areas{
		fmt.Println(area.Name)
	}

	return err
}

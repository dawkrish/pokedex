package commands

import(
	"fmt"
	"net/http"
	"io"
	"errors"
	"encoding/json"

	"github.com/krishnanshagarwal112/pokedex/models"
)

func commandMap(cfg *models.Config) error {
	res, err := http.Get(cfg.Next)
	if err != nil {
		fmt.Println("Error at Get :", err)
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Error at parsing :", err)
		return err
	}

	result := models.Result{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error at decoding :", err)
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

func commandMapB(cfg *models.Config) error {
	if(cfg.Prev == ""){
		err := errors.New("cannot go previous now ")
		fmt.Print(err)
		return err
	}

	res, err := http.Get(cfg.Prev)	

	if err != nil {
		fmt.Println("Error at Get :", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("Error at parsing :", err)
		return err
	}

	result := models.Result{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error at decoding :", err)
		return err
	}

	cfg.Prev= result.Prev
	cfg.Next = result.Next

	areas := result.Results
	for _,area := range areas{
		fmt.Println(area.Name)
	}

	return err
}

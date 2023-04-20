package models


type CliCommand struct {
	Name       string
	Description string
	Callback    func(*Config) error
}

type Config struct{
	Prev string
	Next string
}

type Result struct{
	Prev string `json:"previous"`
	Next string	`json:"next"`
	Results []struct{
		Name string `json:"name"`
		Url string 	`json:"url"`
	} `json:"results"`
}

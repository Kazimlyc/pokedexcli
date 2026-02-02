package main

import (
	"encoding/json"
	"fmt"
)

type LocationAreasResp struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func commandMap(cfg *config, args ...string) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.Next != nil {
		url = *cfg.Next
	}

	data, err := cfg.pokeapiClient.FetchLocations(url)
	if err != nil {
		return err
	}
	var locationData LocationAreasResp
	if err := json.Unmarshal(data, &locationData); err != nil {
		return err
	}
	for i := 0; i < len(locationData.Results); i++ {
		fmt.Println(locationData.Results[i].Name)
	}
	cfg.Next = locationData.Next
	cfg.Previous = locationData.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {

	url := *cfg.Previous

	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	data, err := cfg.pokeapiClient.FetchLocations(url)
	if err != nil {
		return err
	}

	var locationData LocationAreasResp
	if err := json.Unmarshal(data, &locationData); err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		fmt.Println(locationData.Results[i].Name)
	}
	cfg.Next = locationData.Next
	cfg.Previous = locationData.Previous

	return nil
}

package main

import "fmt"

// todo
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	areaName := args[0]

	// TODO: use cfg.pokeapiClient.GetLocation(areaName) and print the results

	_ = areaName

	return nil
}
